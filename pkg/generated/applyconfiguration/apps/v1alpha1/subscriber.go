/*
Copyright The Clusternet Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SubscriberApplyConfiguration represents an declarative configuration of the Subscriber type for use
// with apply.
type SubscriberApplyConfiguration struct {
	ClusterAffinity  *v1.LabelSelector                   `json:"clusterAffinity,omitempty"`
	Weight           *int32                              `json:"weight,omitempty"`
	SubGroupStrategy *SubGroupStrategyApplyConfiguration `json:"subGroupStrategy,omitempty"`
}

// SubscriberApplyConfiguration constructs an declarative configuration of the Subscriber type for use with
// apply.
func Subscriber() *SubscriberApplyConfiguration {
	return &SubscriberApplyConfiguration{}
}

// WithClusterAffinity sets the ClusterAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterAffinity field is set to the value of the last call.
func (b *SubscriberApplyConfiguration) WithClusterAffinity(value v1.LabelSelector) *SubscriberApplyConfiguration {
	b.ClusterAffinity = &value
	return b
}

// WithWeight sets the Weight field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Weight field is set to the value of the last call.
func (b *SubscriberApplyConfiguration) WithWeight(value int32) *SubscriberApplyConfiguration {
	b.Weight = &value
	return b
}

// WithSubGroupStrategy sets the SubGroupStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubGroupStrategy field is set to the value of the last call.
func (b *SubscriberApplyConfiguration) WithSubGroupStrategy(value *SubGroupStrategyApplyConfiguration) *SubscriberApplyConfiguration {
	b.SubGroupStrategy = value
	return b
}
