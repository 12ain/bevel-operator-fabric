/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FabricIdentitySpecApplyConfiguration represents a declarative configuration of the FabricIdentitySpec type for use
// with apply.
type FabricIdentitySpecApplyConfiguration struct {
	Cahost                *string                                            `json:"cahost,omitempty"`
	Caname                *string                                            `json:"caname,omitempty"`
	Caport                *int                                               `json:"caport,omitempty"`
	Catls                 *CatlsApplyConfiguration                           `json:"catls,omitempty"`
	Enrollid              *string                                            `json:"enrollid,omitempty"`
	Enrollsecret          *string                                            `json:"enrollsecret,omitempty"`
	MSPID                 *string                                            `json:"mspid,omitempty"`
	AttributeRequest      []FabricIdentityAttributeRequestApplyConfiguration `json:"attributeRequest,omitempty"`
	Register              *FabricIdentityRegisterApplyConfiguration          `json:"register,omitempty"`
	UpdateCertificateTime *v1.Time                                           `json:"updateCertificateTime,omitempty"`
}

// FabricIdentitySpecApplyConfiguration constructs a declarative configuration of the FabricIdentitySpec type for use with
// apply.
func FabricIdentitySpec() *FabricIdentitySpecApplyConfiguration {
	return &FabricIdentitySpecApplyConfiguration{}
}

// WithCahost sets the Cahost field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cahost field is set to the value of the last call.
func (b *FabricIdentitySpecApplyConfiguration) WithCahost(value string) *FabricIdentitySpecApplyConfiguration {
	b.Cahost = &value
	return b
}

// WithCaname sets the Caname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Caname field is set to the value of the last call.
func (b *FabricIdentitySpecApplyConfiguration) WithCaname(value string) *FabricIdentitySpecApplyConfiguration {
	b.Caname = &value
	return b
}

// WithCaport sets the Caport field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Caport field is set to the value of the last call.
func (b *FabricIdentitySpecApplyConfiguration) WithCaport(value int) *FabricIdentitySpecApplyConfiguration {
	b.Caport = &value
	return b
}

// WithCatls sets the Catls field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Catls field is set to the value of the last call.
func (b *FabricIdentitySpecApplyConfiguration) WithCatls(value *CatlsApplyConfiguration) *FabricIdentitySpecApplyConfiguration {
	b.Catls = value
	return b
}

// WithEnrollid sets the Enrollid field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enrollid field is set to the value of the last call.
func (b *FabricIdentitySpecApplyConfiguration) WithEnrollid(value string) *FabricIdentitySpecApplyConfiguration {
	b.Enrollid = &value
	return b
}

// WithEnrollsecret sets the Enrollsecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enrollsecret field is set to the value of the last call.
func (b *FabricIdentitySpecApplyConfiguration) WithEnrollsecret(value string) *FabricIdentitySpecApplyConfiguration {
	b.Enrollsecret = &value
	return b
}

// WithMSPID sets the MSPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MSPID field is set to the value of the last call.
func (b *FabricIdentitySpecApplyConfiguration) WithMSPID(value string) *FabricIdentitySpecApplyConfiguration {
	b.MSPID = &value
	return b
}

// WithAttributeRequest adds the given value to the AttributeRequest field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AttributeRequest field.
func (b *FabricIdentitySpecApplyConfiguration) WithAttributeRequest(values ...*FabricIdentityAttributeRequestApplyConfiguration) *FabricIdentitySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAttributeRequest")
		}
		b.AttributeRequest = append(b.AttributeRequest, *values[i])
	}
	return b
}

// WithRegister sets the Register field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Register field is set to the value of the last call.
func (b *FabricIdentitySpecApplyConfiguration) WithRegister(value *FabricIdentityRegisterApplyConfiguration) *FabricIdentitySpecApplyConfiguration {
	b.Register = value
	return b
}

// WithUpdateCertificateTime sets the UpdateCertificateTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdateCertificateTime field is set to the value of the last call.
func (b *FabricIdentitySpecApplyConfiguration) WithUpdateCertificateTime(value v1.Time) *FabricIdentitySpecApplyConfiguration {
	b.UpdateCertificateTime = &value
	return b
}
