/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricPeerLoggingApplyConfiguration represents an declarative configuration of the FabricPeerLogging type for use
// with apply.
type FabricPeerLoggingApplyConfiguration struct {
	Level    *string `json:"level,omitempty"`
	Peer     *string `json:"peer,omitempty"`
	Cauthdsl *string `json:"cauthdsl,omitempty"`
	Gossip   *string `json:"gossip,omitempty"`
	Grpc     *string `json:"grpc,omitempty"`
	Ledger   *string `json:"ledger,omitempty"`
	Msp      *string `json:"msp,omitempty"`
	Policies *string `json:"policies,omitempty"`
}

// FabricPeerLoggingApplyConfiguration constructs an declarative configuration of the FabricPeerLogging type for use with
// apply.
func FabricPeerLogging() *FabricPeerLoggingApplyConfiguration {
	return &FabricPeerLoggingApplyConfiguration{}
}

// WithLevel sets the Level field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Level field is set to the value of the last call.
func (b *FabricPeerLoggingApplyConfiguration) WithLevel(value string) *FabricPeerLoggingApplyConfiguration {
	b.Level = &value
	return b
}

// WithPeer sets the Peer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Peer field is set to the value of the last call.
func (b *FabricPeerLoggingApplyConfiguration) WithPeer(value string) *FabricPeerLoggingApplyConfiguration {
	b.Peer = &value
	return b
}

// WithCauthdsl sets the Cauthdsl field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cauthdsl field is set to the value of the last call.
func (b *FabricPeerLoggingApplyConfiguration) WithCauthdsl(value string) *FabricPeerLoggingApplyConfiguration {
	b.Cauthdsl = &value
	return b
}

// WithGossip sets the Gossip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Gossip field is set to the value of the last call.
func (b *FabricPeerLoggingApplyConfiguration) WithGossip(value string) *FabricPeerLoggingApplyConfiguration {
	b.Gossip = &value
	return b
}

// WithGrpc sets the Grpc field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Grpc field is set to the value of the last call.
func (b *FabricPeerLoggingApplyConfiguration) WithGrpc(value string) *FabricPeerLoggingApplyConfiguration {
	b.Grpc = &value
	return b
}

// WithLedger sets the Ledger field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ledger field is set to the value of the last call.
func (b *FabricPeerLoggingApplyConfiguration) WithLedger(value string) *FabricPeerLoggingApplyConfiguration {
	b.Ledger = &value
	return b
}

// WithMsp sets the Msp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Msp field is set to the value of the last call.
func (b *FabricPeerLoggingApplyConfiguration) WithMsp(value string) *FabricPeerLoggingApplyConfiguration {
	b.Msp = &value
	return b
}

// WithPolicies sets the Policies field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Policies field is set to the value of the last call.
func (b *FabricPeerLoggingApplyConfiguration) WithPolicies(value string) *FabricPeerLoggingApplyConfiguration {
	b.Policies = &value
	return b
}
