/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
)

// OrdererServiceApplyConfiguration represents an declarative configuration of the OrdererService type for use
// with apply.
type OrdererServiceApplyConfiguration struct {
	Type *v1alpha1.ServiceType `json:"type,omitempty"`
}

// OrdererServiceApplyConfiguration constructs an declarative configuration of the OrdererService type for use with
// apply.
func OrdererService() *OrdererServiceApplyConfiguration {
	return &OrdererServiceApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *OrdererServiceApplyConfiguration) WithType(value v1alpha1.ServiceType) *OrdererServiceApplyConfiguration {
	b.Type = &value
	return b
}
