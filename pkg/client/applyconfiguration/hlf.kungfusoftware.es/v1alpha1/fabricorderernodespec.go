/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	hlfkungfusoftwareesv1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FabricOrdererNodeSpecApplyConfiguration represents an declarative configuration of the FabricOrdererNodeSpec type for use
// with apply.
type FabricOrdererNodeSpecApplyConfiguration struct {
	Tolerations                 []v1.Toleration                              `json:"tolerations,omitempty"`
	GRPCProxy                   *GRPCProxyApplyConfiguration                 `json:"grpcProxy,omitempty"`
	Affinity                    *v1.Affinity                                 `json:"affinity,omitempty"`
	UpdateCertificateTime       *metav1.Time                                 `json:"updateCertificateTime,omitempty"`
	ServiceMonitor              *ServiceMonitorApplyConfiguration            `json:"serviceMonitor,omitempty"`
	HostAliases                 []v1.HostAlias                               `json:"hostAliases,omitempty"`
	NodeSelector                *v1.NodeSelector                             `json:"nodeSelector,omitempty"`
	Resources                   *v1.ResourceRequirements                     `json:"resources,omitempty"`
	Replicas                    *int                                         `json:"replicas,omitempty"`
	Image                       *string                                      `json:"image,omitempty"`
	Tag                         *string                                      `json:"tag,omitempty"`
	PullPolicy                  *v1.PullPolicy                               `json:"pullPolicy,omitempty"`
	MspID                       *string                                      `json:"mspID,omitempty"`
	ImagePullSecrets            []v1.LocalObjectReference                    `json:"imagePullSecrets,omitempty"`
	Genesis                     *string                                      `json:"genesis,omitempty"`
	BootstrapMethod             *hlfkungfusoftwareesv1alpha1.BootstrapMethod `json:"bootstrapMethod,omitempty"`
	ChannelParticipationEnabled *bool                                        `json:"channelParticipationEnabled,omitempty"`
	Storage                     *StorageApplyConfiguration                   `json:"storage,omitempty"`
	Service                     *OrdererNodeServiceApplyConfiguration        `json:"service,omitempty"`
	Secret                      *SecretApplyConfiguration                    `json:"secret,omitempty"`
	Istio                       *FabricIstioApplyConfiguration               `json:"istio,omitempty"`
	AdminIstio                  *FabricIstioApplyConfiguration               `json:"adminIstio,omitempty"`
	Env                         []v1.EnvVar                                  `json:"env,omitempty"`
}

// FabricOrdererNodeSpecApplyConfiguration constructs an declarative configuration of the FabricOrdererNodeSpec type for use with
// apply.
func FabricOrdererNodeSpec() *FabricOrdererNodeSpecApplyConfiguration {
	return &FabricOrdererNodeSpecApplyConfiguration{}
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithTolerations(values ...v1.Toleration) *FabricOrdererNodeSpecApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithGRPCProxy sets the GRPCProxy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GRPCProxy field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithGRPCProxy(value *GRPCProxyApplyConfiguration) *FabricOrdererNodeSpecApplyConfiguration {
	b.GRPCProxy = value
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithAffinity(value v1.Affinity) *FabricOrdererNodeSpecApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithUpdateCertificateTime sets the UpdateCertificateTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdateCertificateTime field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithUpdateCertificateTime(value metav1.Time) *FabricOrdererNodeSpecApplyConfiguration {
	b.UpdateCertificateTime = &value
	return b
}

// WithServiceMonitor sets the ServiceMonitor field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceMonitor field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithServiceMonitor(value *ServiceMonitorApplyConfiguration) *FabricOrdererNodeSpecApplyConfiguration {
	b.ServiceMonitor = value
	return b
}

// WithHostAliases adds the given value to the HostAliases field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HostAliases field.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithHostAliases(values ...v1.HostAlias) *FabricOrdererNodeSpecApplyConfiguration {
	for i := range values {
		b.HostAliases = append(b.HostAliases, values[i])
	}
	return b
}

// WithNodeSelector sets the NodeSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeSelector field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithNodeSelector(value v1.NodeSelector) *FabricOrdererNodeSpecApplyConfiguration {
	b.NodeSelector = &value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithResources(value v1.ResourceRequirements) *FabricOrdererNodeSpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithReplicas(value int) *FabricOrdererNodeSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithImage(value string) *FabricOrdererNodeSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithTag sets the Tag field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tag field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithTag(value string) *FabricOrdererNodeSpecApplyConfiguration {
	b.Tag = &value
	return b
}

// WithPullPolicy sets the PullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PullPolicy field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithPullPolicy(value v1.PullPolicy) *FabricOrdererNodeSpecApplyConfiguration {
	b.PullPolicy = &value
	return b
}

// WithMspID sets the MspID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MspID field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithMspID(value string) *FabricOrdererNodeSpecApplyConfiguration {
	b.MspID = &value
	return b
}

// WithImagePullSecrets adds the given value to the ImagePullSecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImagePullSecrets field.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithImagePullSecrets(values ...v1.LocalObjectReference) *FabricOrdererNodeSpecApplyConfiguration {
	for i := range values {
		b.ImagePullSecrets = append(b.ImagePullSecrets, values[i])
	}
	return b
}

// WithGenesis sets the Genesis field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Genesis field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithGenesis(value string) *FabricOrdererNodeSpecApplyConfiguration {
	b.Genesis = &value
	return b
}

// WithBootstrapMethod sets the BootstrapMethod field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BootstrapMethod field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithBootstrapMethod(value hlfkungfusoftwareesv1alpha1.BootstrapMethod) *FabricOrdererNodeSpecApplyConfiguration {
	b.BootstrapMethod = &value
	return b
}

// WithChannelParticipationEnabled sets the ChannelParticipationEnabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ChannelParticipationEnabled field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithChannelParticipationEnabled(value bool) *FabricOrdererNodeSpecApplyConfiguration {
	b.ChannelParticipationEnabled = &value
	return b
}

// WithStorage sets the Storage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Storage field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithStorage(value *StorageApplyConfiguration) *FabricOrdererNodeSpecApplyConfiguration {
	b.Storage = value
	return b
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithService(value *OrdererNodeServiceApplyConfiguration) *FabricOrdererNodeSpecApplyConfiguration {
	b.Service = value
	return b
}

// WithSecret sets the Secret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Secret field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithSecret(value *SecretApplyConfiguration) *FabricOrdererNodeSpecApplyConfiguration {
	b.Secret = value
	return b
}

// WithIstio sets the Istio field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Istio field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithIstio(value *FabricIstioApplyConfiguration) *FabricOrdererNodeSpecApplyConfiguration {
	b.Istio = value
	return b
}

// WithAdminIstio sets the AdminIstio field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdminIstio field is set to the value of the last call.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithAdminIstio(value *FabricIstioApplyConfiguration) *FabricOrdererNodeSpecApplyConfiguration {
	b.AdminIstio = value
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *FabricOrdererNodeSpecApplyConfiguration) WithEnv(values ...v1.EnvVar) *FabricOrdererNodeSpecApplyConfiguration {
	for i := range values {
		b.Env = append(b.Env, values[i])
	}
	return b
}
