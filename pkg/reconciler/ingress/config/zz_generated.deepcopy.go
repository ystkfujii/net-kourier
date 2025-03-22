//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	sets "k8s.io/apimachinery/pkg/util/sets"
	pkgconfig "knative.dev/networking/pkg/config"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	if in.Kourier != nil {
		in, out := &in.Kourier, &out.Kourier
		*out = new(Kourier)
		(*in).DeepCopyInto(*out)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(pkgconfig.Config)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kourier) DeepCopyInto(out *Kourier) {
	*out = *in
	if in.CipherSuites != nil {
		in, out := &in.CipherSuites, &out.CipherSuites
		*out = make(sets.Set[string], len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Tracing = in.Tracing
	out.ExternalAuthz = in.ExternalAuthz
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kourier.
func (in *Kourier) DeepCopy() *Kourier {
	if in == nil {
		return nil
	}
	out := new(Kourier)
	in.DeepCopyInto(out)
	return out
}
