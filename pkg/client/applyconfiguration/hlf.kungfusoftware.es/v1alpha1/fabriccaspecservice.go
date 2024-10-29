/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// FabricCASpecServiceApplyConfiguration represents a declarative configuration of the FabricCASpecService type for use
// with apply.
type FabricCASpecServiceApplyConfiguration struct {
	ServiceType *v1.ServiceType `json:"type,omitempty"`
}

// FabricCASpecServiceApplyConfiguration constructs a declarative configuration of the FabricCASpecService type for use with
// apply.
func FabricCASpecService() *FabricCASpecServiceApplyConfiguration {
	return &FabricCASpecServiceApplyConfiguration{}
}

// WithServiceType sets the ServiceType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceType field is set to the value of the last call.
func (b *FabricCASpecServiceApplyConfiguration) WithServiceType(value v1.ServiceType) *FabricCASpecServiceApplyConfiguration {
	b.ServiceType = &value
	return b
}
