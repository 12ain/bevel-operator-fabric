/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// FabricFollowerChannelOrdererApplyConfiguration represents an declarative configuration of the FabricFollowerChannelOrderer type for use
// with apply.
type FabricFollowerChannelOrdererApplyConfiguration struct {
	URL         *string `json:"url,omitempty"`
	Certificate *string `json:"certificate,omitempty"`
}

// FabricFollowerChannelOrdererApplyConfiguration constructs an declarative configuration of the FabricFollowerChannelOrderer type for use with
// apply.
func FabricFollowerChannelOrderer() *FabricFollowerChannelOrdererApplyConfiguration {
	return &FabricFollowerChannelOrdererApplyConfiguration{}
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *FabricFollowerChannelOrdererApplyConfiguration) WithURL(value string) *FabricFollowerChannelOrdererApplyConfiguration {
	b.URL = &value
	return b
}

// WithCertificate sets the Certificate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Certificate field is set to the value of the last call.
func (b *FabricFollowerChannelOrdererApplyConfiguration) WithCertificate(value string) *FabricFollowerChannelOrdererApplyConfiguration {
	b.Certificate = &value
	return b
}
