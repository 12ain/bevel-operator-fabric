/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	status "github.com/kfsoftware/hlf-operator/pkg/status"
)

// FabricOperationsConsoleStatusApplyConfiguration represents an declarative configuration of the FabricOperationsConsoleStatus type for use
// with apply.
type FabricOperationsConsoleStatusApplyConfiguration struct {
	Conditions *status.Conditions         `json:"conditions,omitempty"`
	Message    *string                    `json:"message,omitempty"`
	Status     *v1alpha1.DeploymentStatus `json:"status,omitempty"`
}

// FabricOperationsConsoleStatusApplyConfiguration constructs an declarative configuration of the FabricOperationsConsoleStatus type for use with
// apply.
func FabricOperationsConsoleStatus() *FabricOperationsConsoleStatusApplyConfiguration {
	return &FabricOperationsConsoleStatusApplyConfiguration{}
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *FabricOperationsConsoleStatusApplyConfiguration) WithConditions(value status.Conditions) *FabricOperationsConsoleStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *FabricOperationsConsoleStatusApplyConfiguration) WithMessage(value string) *FabricOperationsConsoleStatusApplyConfiguration {
	b.Message = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *FabricOperationsConsoleStatusApplyConfiguration) WithStatus(value v1alpha1.DeploymentStatus) *FabricOperationsConsoleStatusApplyConfiguration {
	b.Status = &value
	return b
}
