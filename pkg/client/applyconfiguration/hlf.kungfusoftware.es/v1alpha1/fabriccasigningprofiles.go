/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricCASigningProfilesApplyConfiguration represents an declarative configuration of the FabricCASigningProfiles type for use
// with apply.
type FabricCASigningProfilesApplyConfiguration struct {
	CA  *FabricCASigningSignProfileApplyConfiguration `json:"ca,omitempty"`
	TLS *FabricCASigningTLSProfileApplyConfiguration  `json:"tls,omitempty"`
}

// FabricCASigningProfilesApplyConfiguration constructs an declarative configuration of the FabricCASigningProfiles type for use with
// apply.
func FabricCASigningProfiles() *FabricCASigningProfilesApplyConfiguration {
	return &FabricCASigningProfilesApplyConfiguration{}
}

// WithCA sets the CA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CA field is set to the value of the last call.
func (b *FabricCASigningProfilesApplyConfiguration) WithCA(value *FabricCASigningSignProfileApplyConfiguration) *FabricCASigningProfilesApplyConfiguration {
	b.CA = value
	return b
}

// WithTLS sets the TLS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLS field is set to the value of the last call.
func (b *FabricCASigningProfilesApplyConfiguration) WithTLS(value *FabricCASigningTLSProfileApplyConfiguration) *FabricCASigningProfilesApplyConfiguration {
	b.TLS = value
	return b
}
