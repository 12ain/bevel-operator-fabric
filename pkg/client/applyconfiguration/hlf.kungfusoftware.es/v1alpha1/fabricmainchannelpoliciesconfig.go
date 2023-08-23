/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricMainChannelPoliciesConfigApplyConfiguration represents an declarative configuration of the FabricMainChannelPoliciesConfig type for use
// with apply.
type FabricMainChannelPoliciesConfigApplyConfiguration struct {
	Type      *string `json:"type,omitempty"`
	Rule      *string `json:"rule,omitempty"`
	ModPolicy *string `json:"modPolicy,omitempty"`
}

// FabricMainChannelPoliciesConfigApplyConfiguration constructs an declarative configuration of the FabricMainChannelPoliciesConfig type for use with
// apply.
func FabricMainChannelPoliciesConfig() *FabricMainChannelPoliciesConfigApplyConfiguration {
	return &FabricMainChannelPoliciesConfigApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *FabricMainChannelPoliciesConfigApplyConfiguration) WithType(value string) *FabricMainChannelPoliciesConfigApplyConfiguration {
	b.Type = &value
	return b
}

// WithRule sets the Rule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Rule field is set to the value of the last call.
func (b *FabricMainChannelPoliciesConfigApplyConfiguration) WithRule(value string) *FabricMainChannelPoliciesConfigApplyConfiguration {
	b.Rule = &value
	return b
}

// WithModPolicy sets the ModPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ModPolicy field is set to the value of the last call.
func (b *FabricMainChannelPoliciesConfigApplyConfiguration) WithModPolicy(value string) *FabricMainChannelPoliciesConfigApplyConfiguration {
	b.ModPolicy = &value
	return b
}
