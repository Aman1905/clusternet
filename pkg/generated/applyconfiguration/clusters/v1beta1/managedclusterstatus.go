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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ManagedClusterStatusApplyConfiguration represents an declarative configuration of the ManagedClusterStatus type for use
// with apply.
type ManagedClusterStatusApplyConfiguration struct {
	LastObservedTime          *v1.Time                          `json:"lastObservedTime,omitempty"`
	KubernetesVersion         *string                           `json:"k8sVersion,omitempty"`
	Platform                  *string                           `json:"platform,omitempty"`
	APIServerURL              *string                           `json:"apiserverURL,omitempty"`
	Healthz                   *bool                             `json:"healthz,omitempty"`
	Livez                     *bool                             `json:"livez,omitempty"`
	Readyz                    *bool                             `json:"readyz,omitempty"`
	AppPusher                 *bool                             `json:"appPusher,omitempty"`
	UseSocket                 *bool                             `json:"useSocket,omitempty"`
	Allocatable               *corev1.ResourceList              `json:"allocatable,omitempty"`
	Capacity                  *corev1.ResourceList              `json:"capacity,omitempty"`
	ClusterCIDR               *string                           `json:"clusterCIDR,omitempty"`
	ServiceCIDR               *string                           `json:"serviceCIDR,omitempty"`
	NodeStatistics            *NodeStatisticsApplyConfiguration `json:"nodeStatistics,omitempty"`
	PodStatistics             *PodStatisticsApplyConfiguration  `json:"podStatistics,omitempty"`
	ResourceUsage             *ResourceUsageApplyConfiguration  `json:"resourceUsage,omitempty"`
	Conditions                []v1.Condition                    `json:"conditions,omitempty"`
	HeartbeatFrequencySeconds *int64                            `json:"heartbeatFrequencySeconds,omitempty"`
	PredictorEnabled          *bool                             `json:"predictorEnabled,omitempty"`
	PredictorAddress          *string                           `json:"predictorAddress,omitempty"`
	PredictorDirectAccess     *bool                             `json:"predictorDirectAccess,omitempty"`
}

// ManagedClusterStatusApplyConfiguration constructs an declarative configuration of the ManagedClusterStatus type for use with
// apply.
func ManagedClusterStatus() *ManagedClusterStatusApplyConfiguration {
	return &ManagedClusterStatusApplyConfiguration{}
}

// WithLastObservedTime sets the LastObservedTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastObservedTime field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithLastObservedTime(value v1.Time) *ManagedClusterStatusApplyConfiguration {
	b.LastObservedTime = &value
	return b
}

// WithKubernetesVersion sets the KubernetesVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KubernetesVersion field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithKubernetesVersion(value string) *ManagedClusterStatusApplyConfiguration {
	b.KubernetesVersion = &value
	return b
}

// WithPlatform sets the Platform field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Platform field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithPlatform(value string) *ManagedClusterStatusApplyConfiguration {
	b.Platform = &value
	return b
}

// WithAPIServerURL sets the APIServerURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIServerURL field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithAPIServerURL(value string) *ManagedClusterStatusApplyConfiguration {
	b.APIServerURL = &value
	return b
}

// WithHealthz sets the Healthz field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Healthz field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithHealthz(value bool) *ManagedClusterStatusApplyConfiguration {
	b.Healthz = &value
	return b
}

// WithLivez sets the Livez field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Livez field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithLivez(value bool) *ManagedClusterStatusApplyConfiguration {
	b.Livez = &value
	return b
}

// WithReadyz sets the Readyz field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Readyz field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithReadyz(value bool) *ManagedClusterStatusApplyConfiguration {
	b.Readyz = &value
	return b
}

// WithAppPusher sets the AppPusher field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AppPusher field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithAppPusher(value bool) *ManagedClusterStatusApplyConfiguration {
	b.AppPusher = &value
	return b
}

// WithUseSocket sets the UseSocket field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UseSocket field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithUseSocket(value bool) *ManagedClusterStatusApplyConfiguration {
	b.UseSocket = &value
	return b
}

// WithAllocatable sets the Allocatable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Allocatable field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithAllocatable(value corev1.ResourceList) *ManagedClusterStatusApplyConfiguration {
	b.Allocatable = &value
	return b
}

// WithCapacity sets the Capacity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Capacity field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithCapacity(value corev1.ResourceList) *ManagedClusterStatusApplyConfiguration {
	b.Capacity = &value
	return b
}

// WithClusterCIDR sets the ClusterCIDR field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterCIDR field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithClusterCIDR(value string) *ManagedClusterStatusApplyConfiguration {
	b.ClusterCIDR = &value
	return b
}

// WithServiceCIDR sets the ServiceCIDR field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceCIDR field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithServiceCIDR(value string) *ManagedClusterStatusApplyConfiguration {
	b.ServiceCIDR = &value
	return b
}

// WithNodeStatistics sets the NodeStatistics field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeStatistics field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithNodeStatistics(value *NodeStatisticsApplyConfiguration) *ManagedClusterStatusApplyConfiguration {
	b.NodeStatistics = value
	return b
}

// WithPodStatistics sets the PodStatistics field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodStatistics field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithPodStatistics(value *PodStatisticsApplyConfiguration) *ManagedClusterStatusApplyConfiguration {
	b.PodStatistics = value
	return b
}

// WithResourceUsage sets the ResourceUsage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceUsage field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithResourceUsage(value *ResourceUsageApplyConfiguration) *ManagedClusterStatusApplyConfiguration {
	b.ResourceUsage = value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ManagedClusterStatusApplyConfiguration) WithConditions(values ...v1.Condition) *ManagedClusterStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}

// WithHeartbeatFrequencySeconds sets the HeartbeatFrequencySeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HeartbeatFrequencySeconds field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithHeartbeatFrequencySeconds(value int64) *ManagedClusterStatusApplyConfiguration {
	b.HeartbeatFrequencySeconds = &value
	return b
}

// WithPredictorEnabled sets the PredictorEnabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PredictorEnabled field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithPredictorEnabled(value bool) *ManagedClusterStatusApplyConfiguration {
	b.PredictorEnabled = &value
	return b
}

// WithPredictorAddress sets the PredictorAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PredictorAddress field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithPredictorAddress(value string) *ManagedClusterStatusApplyConfiguration {
	b.PredictorAddress = &value
	return b
}

// WithPredictorDirectAccess sets the PredictorDirectAccess field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PredictorDirectAccess field is set to the value of the last call.
func (b *ManagedClusterStatusApplyConfiguration) WithPredictorDirectAccess(value bool) *ManagedClusterStatusApplyConfiguration {
	b.PredictorDirectAccess = &value
	return b
}
