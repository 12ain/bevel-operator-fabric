/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCANamesApplyConfiguration represents an declarative configuration of the FabricCANames type for use
// with apply.
type FabricCANamesApplyConfiguration struct {
	C  *string `json:"C,omitempty"`
	ST *string `json:"ST,omitempty"`
	O  *string `json:"O,omitempty"`
	L  *string `json:"L,omitempty"`
	OU *string `json:"OU,omitempty"`
}

// FabricCANamesApplyConfiguration constructs an declarative configuration of the FabricCANames type for use with
// apply.
func FabricCANames() *FabricCANamesApplyConfiguration {
	return &FabricCANamesApplyConfiguration{}
}

// WithC sets the C field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the C field is set to the value of the last call.
func (b *FabricCANamesApplyConfiguration) WithC(value string) *FabricCANamesApplyConfiguration {
	b.C = &value
	return b
}

// WithST sets the ST field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ST field is set to the value of the last call.
func (b *FabricCANamesApplyConfiguration) WithST(value string) *FabricCANamesApplyConfiguration {
	b.ST = &value
	return b
}

// WithO sets the O field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the O field is set to the value of the last call.
func (b *FabricCANamesApplyConfiguration) WithO(value string) *FabricCANamesApplyConfiguration {
	b.O = &value
	return b
}

// WithL sets the L field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the L field is set to the value of the last call.
func (b *FabricCANamesApplyConfiguration) WithL(value string) *FabricCANamesApplyConfiguration {
	b.L = &value
	return b
}

// WithOU sets the OU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OU field is set to the value of the last call.
func (b *FabricCANamesApplyConfiguration) WithOU(value string) *FabricCANamesApplyConfiguration {
	b.OU = &value
	return b
}
