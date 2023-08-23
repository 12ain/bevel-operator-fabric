/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelConsenterApplyConfiguration represents an declarative configuration of the FabricMainChannelConsenter type for use
// with apply.
type FabricMainChannelConsenterApplyConfiguration struct {
	Host    *string `json:"host,omitempty"`
	Port    *int    `json:"port,omitempty"`
	TLSCert *string `json:"tlsCert,omitempty"`
}

// FabricMainChannelConsenterApplyConfiguration constructs an declarative configuration of the FabricMainChannelConsenter type for use with
// apply.
func FabricMainChannelConsenter() *FabricMainChannelConsenterApplyConfiguration {
	return &FabricMainChannelConsenterApplyConfiguration{}
}

// WithHost sets the Host field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Host field is set to the value of the last call.
func (b *FabricMainChannelConsenterApplyConfiguration) WithHost(value string) *FabricMainChannelConsenterApplyConfiguration {
	b.Host = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *FabricMainChannelConsenterApplyConfiguration) WithPort(value int) *FabricMainChannelConsenterApplyConfiguration {
	b.Port = &value
	return b
}

// WithTLSCert sets the TLSCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSCert field is set to the value of the last call.
func (b *FabricMainChannelConsenterApplyConfiguration) WithTLSCert(value string) *FabricMainChannelConsenterApplyConfiguration {
	b.TLSCert = &value
	return b
}
