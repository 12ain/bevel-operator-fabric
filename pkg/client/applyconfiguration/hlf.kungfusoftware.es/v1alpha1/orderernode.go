/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// OrdererNodeApplyConfiguration represents an declarative configuration of the OrdererNode type for use
// with apply.
type OrdererNodeApplyConfiguration struct {
	ID         *string                                  `json:"id,omitempty"`
	Host       *string                                  `json:"host,omitempty"`
	Port       *int                                     `json:"port,omitempty"`
	Enrollment *OrdererNodeEnrollmentApplyConfiguration `json:"enrollment,omitempty"`
}

// OrdererNodeApplyConfiguration constructs an declarative configuration of the OrdererNode type for use with
// apply.
func OrdererNode() *OrdererNodeApplyConfiguration {
	return &OrdererNodeApplyConfiguration{}
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *OrdererNodeApplyConfiguration) WithID(value string) *OrdererNodeApplyConfiguration {
	b.ID = &value
	return b
}

// WithHost sets the Host field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Host field is set to the value of the last call.
func (b *OrdererNodeApplyConfiguration) WithHost(value string) *OrdererNodeApplyConfiguration {
	b.Host = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *OrdererNodeApplyConfiguration) WithPort(value int) *OrdererNodeApplyConfiguration {
	b.Port = &value
	return b
}

// WithEnrollment sets the Enrollment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enrollment field is set to the value of the last call.
func (b *OrdererNodeApplyConfiguration) WithEnrollment(value *OrdererNodeEnrollmentApplyConfiguration) *OrdererNodeApplyConfiguration {
	b.Enrollment = value
	return b
}
