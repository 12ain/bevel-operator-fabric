/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// FabricPeerCouchdbExporterApplyConfiguration represents an declarative configuration of the FabricPeerCouchdbExporter type for use
// with apply.
type FabricPeerCouchdbExporterApplyConfiguration struct {
	Enabled         *bool          `json:"enabled,omitempty"`
	Image           *string        `json:"image,omitempty"`
	Tag             *string        `json:"tag,omitempty"`
	ImagePullPolicy *v1.PullPolicy `json:"imagePullPolicy,omitempty"`
}

// FabricPeerCouchdbExporterApplyConfiguration constructs an declarative configuration of the FabricPeerCouchdbExporter type for use with
// apply.
func FabricPeerCouchdbExporter() *FabricPeerCouchdbExporterApplyConfiguration {
	return &FabricPeerCouchdbExporterApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *FabricPeerCouchdbExporterApplyConfiguration) WithEnabled(value bool) *FabricPeerCouchdbExporterApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *FabricPeerCouchdbExporterApplyConfiguration) WithImage(value string) *FabricPeerCouchdbExporterApplyConfiguration {
	b.Image = &value
	return b
}

// WithTag sets the Tag field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tag field is set to the value of the last call.
func (b *FabricPeerCouchdbExporterApplyConfiguration) WithTag(value string) *FabricPeerCouchdbExporterApplyConfiguration {
	b.Tag = &value
	return b
}

// WithImagePullPolicy sets the ImagePullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImagePullPolicy field is set to the value of the last call.
func (b *FabricPeerCouchdbExporterApplyConfiguration) WithImagePullPolicy(value v1.PullPolicy) *FabricPeerCouchdbExporterApplyConfiguration {
	b.ImagePullPolicy = &value
	return b
}
