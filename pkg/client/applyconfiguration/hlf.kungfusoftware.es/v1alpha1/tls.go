/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

// TLSApplyConfiguration represents an declarative configuration of the TLS type for use
// with apply.
type TLSApplyConfiguration struct {
	Cahost       *string                                `json:"cahost,omitempty"`
	Caname       *string                                `json:"caname,omitempty"`
	Caport       *int                                   `json:"caport,omitempty"`
	Catls        *CatlsApplyConfiguration               `json:"catls,omitempty"`
	Csr          *CsrApplyConfiguration                 `json:"csr,omitempty"`
	Enrollid     *string                                `json:"enrollid,omitempty"`
	Enrollsecret *string                                `json:"enrollsecret,omitempty"`
	External     *ExternalCertificateApplyConfiguration `json:"external,omitempty"`
}

// TLSApplyConfiguration constructs an declarative configuration of the TLS type for use with
// apply.
func TLS() *TLSApplyConfiguration {
	return &TLSApplyConfiguration{}
}

// WithCahost sets the Cahost field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cahost field is set to the value of the last call.
func (b *TLSApplyConfiguration) WithCahost(value string) *TLSApplyConfiguration {
	b.Cahost = &value
	return b
}

// WithCaname sets the Caname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Caname field is set to the value of the last call.
func (b *TLSApplyConfiguration) WithCaname(value string) *TLSApplyConfiguration {
	b.Caname = &value
	return b
}

// WithCaport sets the Caport field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Caport field is set to the value of the last call.
func (b *TLSApplyConfiguration) WithCaport(value int) *TLSApplyConfiguration {
	b.Caport = &value
	return b
}

// WithCatls sets the Catls field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Catls field is set to the value of the last call.
func (b *TLSApplyConfiguration) WithCatls(value *CatlsApplyConfiguration) *TLSApplyConfiguration {
	b.Catls = value
	return b
}

// WithCsr sets the Csr field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Csr field is set to the value of the last call.
func (b *TLSApplyConfiguration) WithCsr(value *CsrApplyConfiguration) *TLSApplyConfiguration {
	b.Csr = value
	return b
}

// WithEnrollid sets the Enrollid field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enrollid field is set to the value of the last call.
func (b *TLSApplyConfiguration) WithEnrollid(value string) *TLSApplyConfiguration {
	b.Enrollid = &value
	return b
}

// WithEnrollsecret sets the Enrollsecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enrollsecret field is set to the value of the last call.
func (b *TLSApplyConfiguration) WithEnrollsecret(value string) *TLSApplyConfiguration {
	b.Enrollsecret = &value
	return b
}

// WithExternal sets the External field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the External field is set to the value of the last call.
func (b *TLSApplyConfiguration) WithExternal(value *ExternalCertificateApplyConfiguration) *TLSApplyConfiguration {
	b.External = value
	return b
}
