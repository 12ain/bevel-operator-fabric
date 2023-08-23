/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// EnrollmentApplyConfiguration represents an declarative configuration of the Enrollment type for use
// with apply.
type EnrollmentApplyConfiguration struct {
	Component *ComponentApplyConfiguration `json:"component,omitempty"`
	TLS       *TLSApplyConfiguration       `json:"tls,omitempty"`
}

// EnrollmentApplyConfiguration constructs an declarative configuration of the Enrollment type for use with
// apply.
func Enrollment() *EnrollmentApplyConfiguration {
	return &EnrollmentApplyConfiguration{}
}

// WithComponent sets the Component field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Component field is set to the value of the last call.
func (b *EnrollmentApplyConfiguration) WithComponent(value *ComponentApplyConfiguration) *EnrollmentApplyConfiguration {
	b.Component = value
	return b
}

// WithTLS sets the TLS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLS field is set to the value of the last call.
func (b *EnrollmentApplyConfiguration) WithTLS(value *TLSApplyConfiguration) *EnrollmentApplyConfiguration {
	b.TLS = value
	return b
}
