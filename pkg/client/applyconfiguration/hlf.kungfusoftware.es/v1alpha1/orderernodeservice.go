/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// OrdererNodeServiceApplyConfiguration represents an declarative configuration of the OrdererNodeService type for use
// with apply.
type OrdererNodeServiceApplyConfiguration struct {
	Type               *v1.ServiceType `json:"type,omitempty"`
	NodePortOperations *int            `json:"nodePortOperations,omitempty"`
	NodePortRequest    *int            `json:"nodePortRequest,omitempty"`
}

// OrdererNodeServiceApplyConfiguration constructs an declarative configuration of the OrdererNodeService type for use with
// apply.
func OrdererNodeService() *OrdererNodeServiceApplyConfiguration {
	return &OrdererNodeServiceApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *OrdererNodeServiceApplyConfiguration) WithType(value v1.ServiceType) *OrdererNodeServiceApplyConfiguration {
	b.Type = &value
	return b
}

// WithNodePortOperations sets the NodePortOperations field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodePortOperations field is set to the value of the last call.
func (b *OrdererNodeServiceApplyConfiguration) WithNodePortOperations(value int) *OrdererNodeServiceApplyConfiguration {
	b.NodePortOperations = &value
	return b
}

// WithNodePortRequest sets the NodePortRequest field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodePortRequest field is set to the value of the last call.
func (b *OrdererNodeServiceApplyConfiguration) WithNodePortRequest(value int) *OrdererNodeServiceApplyConfiguration {
	b.NodePortRequest = &value
	return b
}
