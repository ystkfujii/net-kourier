/*
Copyright 2021 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"context"

	network "knative.dev/networking/pkg"
	netconfig "knative.dev/networking/pkg/config"
	"knative.dev/pkg/configmap"
)

type cfgKey struct{}

// Config contains the configmaps requires for revision reconciliation.
// +k8s:deepcopy-gen=true
type Config struct {
	Kourier *Kourier
	Network *netconfig.Config
}

// FromContext loads the configuration from the context.
func FromContext(ctx context.Context) *Config {
	return ctx.Value(cfgKey{}).(*Config)
}

func FromContextOrDefaults(ctx context.Context) *Config {
	if cfg, ok := ctx.Value(cfgKey{}).(*Config); ok {
		return cfg
	}
	return &Config{
		Kourier: defaultKourierConfig(),
		Network: defaultNetworkConfig(),
	}
}

func defaultNetworkConfig() *netconfig.Config {
	return &netconfig.Config{
		SystemInternalTLS: netconfig.EncryptionDisabled,
	}
}

// ToContext persists the configuration to the context.
func ToContext(ctx context.Context, c *Config) context.Context {
	return context.WithValue(ctx, cfgKey{}, c)
}

// Store is a typed wrapper around configmap.UntypedStore to handle our configmaps.
// +k8s:deepcopy-gen=false
type Store struct {
	*configmap.UntypedStore
}

// NewStore creates a new store of Configs and optionally calls functions when ConfigMaps are updated for Revisions
func NewStore(logger configmap.Logger, onAfterStore ...func(name string, value interface{})) *Store {
	store := &Store{
		UntypedStore: configmap.NewUntypedStore(
			"kourier",
			logger,
			configmap.Constructors{
				ConfigName:              NewKourierConfigFromConfigMap,
				netconfig.ConfigMapName: network.NewConfigFromConfigMap,
			},
			onAfterStore...,
		),
	}
	return store
}

// ToContext persists the config on the context.
func (s *Store) ToContext(ctx context.Context) context.Context {
	return ToContext(ctx, s.Load())
}

// Load returns the config from the store.
func (s *Store) Load() *Config {
	return &Config{
		Kourier: s.UntypedLoad(ConfigName).(*Kourier).DeepCopy(),
		Network: s.UntypedLoad(netconfig.ConfigMapName).(*netconfig.Config).DeepCopy(),
	}
}
