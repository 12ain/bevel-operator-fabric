/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricFollowerChannelAnchorPeerApplyConfiguration represents an declarative configuration of the FabricFollowerChannelAnchorPeer type for use
// with apply.
type FabricFollowerChannelAnchorPeerApplyConfiguration struct {
	Host *string `json:"host,omitempty"`
	Port *int    `json:"port,omitempty"`
}

// FabricFollowerChannelAnchorPeerApplyConfiguration constructs an declarative configuration of the FabricFollowerChannelAnchorPeer type for use with
// apply.
func FabricFollowerChannelAnchorPeer() *FabricFollowerChannelAnchorPeerApplyConfiguration {
	return &FabricFollowerChannelAnchorPeerApplyConfiguration{}
}

// WithHost sets the Host field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Host field is set to the value of the last call.
func (b *FabricFollowerChannelAnchorPeerApplyConfiguration) WithHost(value string) *FabricFollowerChannelAnchorPeerApplyConfiguration {
	b.Host = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *FabricFollowerChannelAnchorPeerApplyConfiguration) WithPort(value int) *FabricFollowerChannelAnchorPeerApplyConfiguration {
	b.Port = &value
	return b
}
