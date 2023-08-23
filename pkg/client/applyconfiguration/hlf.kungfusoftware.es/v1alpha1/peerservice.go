/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// PeerServiceApplyConfiguration represents an declarative configuration of the PeerService type for use
// with apply.
type PeerServiceApplyConfiguration struct {
	Type *v1.ServiceType `json:"type,omitempty"`
}

// PeerServiceApplyConfiguration constructs an declarative configuration of the PeerService type for use with
// apply.
func PeerService() *PeerServiceApplyConfiguration {
	return &PeerServiceApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *PeerServiceApplyConfiguration) WithType(value v1.ServiceType) *PeerServiceApplyConfiguration {
	b.Type = &value
	return b
}
