/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// ExternalBuilderApplyConfiguration represents an declarative configuration of the ExternalBuilder type for use
// with apply.
type ExternalBuilderApplyConfiguration struct {
	Name                 *string  `json:"name,omitempty"`
	Path                 *string  `json:"path,omitempty"`
	PropagateEnvironment []string `json:"propagateEnvironment,omitempty"`
}

// ExternalBuilderApplyConfiguration constructs an declarative configuration of the ExternalBuilder type for use with
// apply.
func ExternalBuilder() *ExternalBuilderApplyConfiguration {
	return &ExternalBuilderApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ExternalBuilderApplyConfiguration) WithName(value string) *ExternalBuilderApplyConfiguration {
	b.Name = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *ExternalBuilderApplyConfiguration) WithPath(value string) *ExternalBuilderApplyConfiguration {
	b.Path = &value
	return b
}

// WithPropagateEnvironment adds the given value to the PropagateEnvironment field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PropagateEnvironment field.
func (b *ExternalBuilderApplyConfiguration) WithPropagateEnvironment(values ...string) *ExternalBuilderApplyConfiguration {
	for i := range values {
		b.PropagateEnvironment = append(b.PropagateEnvironment, values[i])
	}
	return b
}
