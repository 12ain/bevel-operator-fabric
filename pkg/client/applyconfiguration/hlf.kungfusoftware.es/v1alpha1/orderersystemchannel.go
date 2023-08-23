/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// OrdererSystemChannelApplyConfiguration represents an declarative configuration of the OrdererSystemChannel type for use
// with apply.
type OrdererSystemChannelApplyConfiguration struct {
	Name   *string                          `json:"name,omitempty"`
	Config *ChannelConfigApplyConfiguration `json:"config,omitempty"`
}

// OrdererSystemChannelApplyConfiguration constructs an declarative configuration of the OrdererSystemChannel type for use with
// apply.
func OrdererSystemChannel() *OrdererSystemChannelApplyConfiguration {
	return &OrdererSystemChannelApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *OrdererSystemChannelApplyConfiguration) WithName(value string) *OrdererSystemChannelApplyConfiguration {
	b.Name = &value
	return b
}

// WithConfig sets the Config field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Config field is set to the value of the last call.
func (b *OrdererSystemChannelApplyConfiguration) WithConfig(value *ChannelConfigApplyConfiguration) *OrdererSystemChannelApplyConfiguration {
	b.Config = value
	return b
}
