/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricChaincodeApproveSpecApplyConfiguration represents a declarative configuration of the FabricChaincodeApproveSpec type for use
// with apply.
type FabricChaincodeApproveSpecApplyConfiguration struct {
	ChaincodeName          *string                                      `json:"chaincodeName,omitempty"`
	ChannelName            *string                                      `json:"channelName,omitempty"`
	InitRequired           *bool                                        `json:"initRequired,omitempty"`
	MSPID                  *string                                      `json:"mspID,omitempty"`
	PackageID              *string                                      `json:"packageId,omitempty"`
	Version                *string                                      `json:"version,omitempty"`
	Sequence               *int64                                       `json:"sequence,omitempty"`
	EndorsementPolicy      *string                                      `json:"endorsementPolicy,omitempty"`
	PrivateDataCollections []PrivateDataCollectionApplyConfiguration    `json:"pdc,omitempty"`
	HLFIdentity            *HLFIdentityApplyConfiguration               `json:"hlfIdentity,omitempty"`
	Peers                  []FabricPeerInternalRefApplyConfiguration    `json:"peers,omitempty"`
	ExternalPeers          []FabricPeerExternalRefApplyConfiguration    `json:"externalPeers,omitempty"`
	Orderers               []FabricOrdererInternalRefApplyConfiguration `json:"orderers,omitempty"`
	ExternalOrderers       []FabricOrdererExternalRefApplyConfiguration `json:"externalOrderers,omitempty"`
}

// FabricChaincodeApproveSpecApplyConfiguration constructs a declarative configuration of the FabricChaincodeApproveSpec type for use with
// apply.
func FabricChaincodeApproveSpec() *FabricChaincodeApproveSpecApplyConfiguration {
	return &FabricChaincodeApproveSpecApplyConfiguration{}
}

// WithChaincodeName sets the ChaincodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ChaincodeName field is set to the value of the last call.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithChaincodeName(value string) *FabricChaincodeApproveSpecApplyConfiguration {
	b.ChaincodeName = &value
	return b
}

// WithChannelName sets the ChannelName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ChannelName field is set to the value of the last call.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithChannelName(value string) *FabricChaincodeApproveSpecApplyConfiguration {
	b.ChannelName = &value
	return b
}

// WithInitRequired sets the InitRequired field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InitRequired field is set to the value of the last call.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithInitRequired(value bool) *FabricChaincodeApproveSpecApplyConfiguration {
	b.InitRequired = &value
	return b
}

// WithMSPID sets the MSPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MSPID field is set to the value of the last call.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithMSPID(value string) *FabricChaincodeApproveSpecApplyConfiguration {
	b.MSPID = &value
	return b
}

// WithPackageID sets the PackageID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PackageID field is set to the value of the last call.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithPackageID(value string) *FabricChaincodeApproveSpecApplyConfiguration {
	b.PackageID = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithVersion(value string) *FabricChaincodeApproveSpecApplyConfiguration {
	b.Version = &value
	return b
}

// WithSequence sets the Sequence field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Sequence field is set to the value of the last call.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithSequence(value int64) *FabricChaincodeApproveSpecApplyConfiguration {
	b.Sequence = &value
	return b
}

// WithEndorsementPolicy sets the EndorsementPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndorsementPolicy field is set to the value of the last call.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithEndorsementPolicy(value string) *FabricChaincodeApproveSpecApplyConfiguration {
	b.EndorsementPolicy = &value
	return b
}

// WithPrivateDataCollections adds the given value to the PrivateDataCollections field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PrivateDataCollections field.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithPrivateDataCollections(values ...*PrivateDataCollectionApplyConfiguration) *FabricChaincodeApproveSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPrivateDataCollections")
		}
		b.PrivateDataCollections = append(b.PrivateDataCollections, *values[i])
	}
	return b
}

// WithHLFIdentity sets the HLFIdentity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HLFIdentity field is set to the value of the last call.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithHLFIdentity(value *HLFIdentityApplyConfiguration) *FabricChaincodeApproveSpecApplyConfiguration {
	b.HLFIdentity = value
	return b
}

// WithPeers adds the given value to the Peers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Peers field.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithPeers(values ...*FabricPeerInternalRefApplyConfiguration) *FabricChaincodeApproveSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPeers")
		}
		b.Peers = append(b.Peers, *values[i])
	}
	return b
}

// WithExternalPeers adds the given value to the ExternalPeers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExternalPeers field.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithExternalPeers(values ...*FabricPeerExternalRefApplyConfiguration) *FabricChaincodeApproveSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExternalPeers")
		}
		b.ExternalPeers = append(b.ExternalPeers, *values[i])
	}
	return b
}

// WithOrderers adds the given value to the Orderers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Orderers field.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithOrderers(values ...*FabricOrdererInternalRefApplyConfiguration) *FabricChaincodeApproveSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOrderers")
		}
		b.Orderers = append(b.Orderers, *values[i])
	}
	return b
}

// WithExternalOrderers adds the given value to the ExternalOrderers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExternalOrderers field.
func (b *FabricChaincodeApproveSpecApplyConfiguration) WithExternalOrderers(values ...*FabricOrdererExternalRefApplyConfiguration) *FabricChaincodeApproveSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExternalOrderers")
		}
		b.ExternalOrderers = append(b.ExternalOrderers, *values[i])
	}
	return b
}