/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// HLFIdentityApplyConfiguration represents an declarative configuration of the HLFIdentity type for use
// with apply.
type HLFIdentityApplyConfiguration struct {
	SecretName      *string `json:"secretName,omitempty"`
	SecretNamespace *string `json:"secretNamespace,omitempty"`
	SecretKey       *string `json:"secretKey,omitempty"`
}

// HLFIdentityApplyConfiguration constructs an declarative configuration of the HLFIdentity type for use with
// apply.
func HLFIdentity() *HLFIdentityApplyConfiguration {
	return &HLFIdentityApplyConfiguration{}
}

// WithSecretName sets the SecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretName field is set to the value of the last call.
func (b *HLFIdentityApplyConfiguration) WithSecretName(value string) *HLFIdentityApplyConfiguration {
	b.SecretName = &value
	return b
}

// WithSecretNamespace sets the SecretNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretNamespace field is set to the value of the last call.
func (b *HLFIdentityApplyConfiguration) WithSecretNamespace(value string) *HLFIdentityApplyConfiguration {
	b.SecretNamespace = &value
	return b
}

// WithSecretKey sets the SecretKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretKey field is set to the value of the last call.
func (b *HLFIdentityApplyConfiguration) WithSecretKey(value string) *HLFIdentityApplyConfiguration {
	b.SecretKey = &value
	return b
}
