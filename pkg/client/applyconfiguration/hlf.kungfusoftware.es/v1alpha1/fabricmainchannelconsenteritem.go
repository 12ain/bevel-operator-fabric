/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricMainChannelConsenterItemApplyConfiguration represents a declarative configuration of the FabricMainChannelConsenterItem type for use
// with apply.
type FabricMainChannelConsenterItemApplyConfiguration struct {
	Id            *uint32 `json:"id,omitempty"`
	Host          *string `json:"host,omitempty"`
	Port          *uint32 `json:"port,omitempty"`
	MspId         *string `json:"msp_id,omitempty"`
	Identity      *string `json:"identity,omitempty"`
	ClientTlsCert *string `json:"client_tls_cert,omitempty"`
	ServerTlsCert *string `json:"server_tls_cert,omitempty"`
}

// FabricMainChannelConsenterItemApplyConfiguration constructs a declarative configuration of the FabricMainChannelConsenterItem type for use with
// apply.
func FabricMainChannelConsenterItem() *FabricMainChannelConsenterItemApplyConfiguration {
	return &FabricMainChannelConsenterItemApplyConfiguration{}
}

// WithId sets the Id field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Id field is set to the value of the last call.
func (b *FabricMainChannelConsenterItemApplyConfiguration) WithId(value uint32) *FabricMainChannelConsenterItemApplyConfiguration {
	b.Id = &value
	return b
}

// WithHost sets the Host field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Host field is set to the value of the last call.
func (b *FabricMainChannelConsenterItemApplyConfiguration) WithHost(value string) *FabricMainChannelConsenterItemApplyConfiguration {
	b.Host = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *FabricMainChannelConsenterItemApplyConfiguration) WithPort(value uint32) *FabricMainChannelConsenterItemApplyConfiguration {
	b.Port = &value
	return b
}

// WithMspId sets the MspId field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MspId field is set to the value of the last call.
func (b *FabricMainChannelConsenterItemApplyConfiguration) WithMspId(value string) *FabricMainChannelConsenterItemApplyConfiguration {
	b.MspId = &value
	return b
}

// WithIdentity sets the Identity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Identity field is set to the value of the last call.
func (b *FabricMainChannelConsenterItemApplyConfiguration) WithIdentity(value string) *FabricMainChannelConsenterItemApplyConfiguration {
	b.Identity = &value
	return b
}

// WithClientTlsCert sets the ClientTlsCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientTlsCert field is set to the value of the last call.
func (b *FabricMainChannelConsenterItemApplyConfiguration) WithClientTlsCert(value string) *FabricMainChannelConsenterItemApplyConfiguration {
	b.ClientTlsCert = &value
	return b
}

// WithServerTlsCert sets the ServerTlsCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServerTlsCert field is set to the value of the last call.
func (b *FabricMainChannelConsenterItemApplyConfiguration) WithServerTlsCert(value string) *FabricMainChannelConsenterItemApplyConfiguration {
	b.ServerTlsCert = &value
	return b
}
