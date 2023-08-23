/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	status "github.com/kfsoftware/hlf-operator/pkg/status"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FabricOrdererNodeStatusApplyConfiguration represents an declarative configuration of the FabricOrdererNodeStatus type for use
// with apply.
type FabricOrdererNodeStatusApplyConfiguration struct {
	Conditions            *status.Conditions         `json:"conditions,omitempty"`
	Status                *v1alpha1.DeploymentStatus `json:"status,omitempty"`
	LastCertificateUpdate *v1.Time                   `json:"lastCertificateUpdate,omitempty"`
	SignCert              *string                    `json:"signCert,omitempty"`
	TlsCert               *string                    `json:"tlsCert,omitempty"`
	SignCACert            *string                    `json:"signCaCert,omitempty"`
	TlsCACert             *string                    `json:"tlsCaCert,omitempty"`
	TlsAdminCert          *string                    `json:"tlsAdminCert,omitempty"`
	OperationsPort        *int                       `json:"operationsPort,omitempty"`
	AdminPort             *int                       `json:"adminPort,omitempty"`
	NodePort              *int                       `json:"port,omitempty"`
	Message               *string                    `json:"message,omitempty"`
}

// FabricOrdererNodeStatusApplyConfiguration constructs an declarative configuration of the FabricOrdererNodeStatus type for use with
// apply.
func FabricOrdererNodeStatus() *FabricOrdererNodeStatusApplyConfiguration {
	return &FabricOrdererNodeStatusApplyConfiguration{}
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithConditions(value status.Conditions) *FabricOrdererNodeStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithStatus(value v1alpha1.DeploymentStatus) *FabricOrdererNodeStatusApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastCertificateUpdate sets the LastCertificateUpdate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastCertificateUpdate field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithLastCertificateUpdate(value v1.Time) *FabricOrdererNodeStatusApplyConfiguration {
	b.LastCertificateUpdate = &value
	return b
}

// WithSignCert sets the SignCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SignCert field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithSignCert(value string) *FabricOrdererNodeStatusApplyConfiguration {
	b.SignCert = &value
	return b
}

// WithTlsCert sets the TlsCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TlsCert field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithTlsCert(value string) *FabricOrdererNodeStatusApplyConfiguration {
	b.TlsCert = &value
	return b
}

// WithSignCACert sets the SignCACert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SignCACert field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithSignCACert(value string) *FabricOrdererNodeStatusApplyConfiguration {
	b.SignCACert = &value
	return b
}

// WithTlsCACert sets the TlsCACert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TlsCACert field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithTlsCACert(value string) *FabricOrdererNodeStatusApplyConfiguration {
	b.TlsCACert = &value
	return b
}

// WithTlsAdminCert sets the TlsAdminCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TlsAdminCert field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithTlsAdminCert(value string) *FabricOrdererNodeStatusApplyConfiguration {
	b.TlsAdminCert = &value
	return b
}

// WithOperationsPort sets the OperationsPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OperationsPort field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithOperationsPort(value int) *FabricOrdererNodeStatusApplyConfiguration {
	b.OperationsPort = &value
	return b
}

// WithAdminPort sets the AdminPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdminPort field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithAdminPort(value int) *FabricOrdererNodeStatusApplyConfiguration {
	b.AdminPort = &value
	return b
}

// WithNodePort sets the NodePort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodePort field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithNodePort(value int) *FabricOrdererNodeStatusApplyConfiguration {
	b.NodePort = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *FabricOrdererNodeStatusApplyConfiguration) WithMessage(value string) *FabricOrdererNodeStatusApplyConfiguration {
	b.Message = &value
	return b
}
