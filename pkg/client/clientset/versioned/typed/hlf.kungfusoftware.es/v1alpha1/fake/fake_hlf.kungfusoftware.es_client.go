/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package fake

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/pkg/client/clientset/versioned/typed/hlf.kungfusoftware.es/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHlfV1alpha1 struct {
	*testing.Fake
}

func (c *FakeHlfV1alpha1) FabricCAs(namespace string) v1alpha1.FabricCAInterface {
	return &FakeFabricCAs{c, namespace}
}

func (c *FakeHlfV1alpha1) FabricChaincodes(namespace string) v1alpha1.FabricChaincodeInterface {
	return &FakeFabricChaincodes{c, namespace}
}

func (c *FakeHlfV1alpha1) FabricExplorers(namespace string) v1alpha1.FabricExplorerInterface {
	return &FakeFabricExplorers{c, namespace}
}

func (c *FakeHlfV1alpha1) FabricFollowerChannels() v1alpha1.FabricFollowerChannelInterface {
	return &FakeFabricFollowerChannels{c}
}

func (c *FakeHlfV1alpha1) FabricIdentities(namespace string) v1alpha1.FabricIdentityInterface {
	return &FakeFabricIdentities{c, namespace}
}

func (c *FakeHlfV1alpha1) FabricMainChannels() v1alpha1.FabricMainChannelInterface {
	return &FakeFabricMainChannels{c}
}

func (c *FakeHlfV1alpha1) FabricNetworkConfigs(namespace string) v1alpha1.FabricNetworkConfigInterface {
	return &FakeFabricNetworkConfigs{c, namespace}
}

func (c *FakeHlfV1alpha1) FabricOperationsConsoles(namespace string) v1alpha1.FabricOperationsConsoleInterface {
	return &FakeFabricOperationsConsoles{c, namespace}
}

func (c *FakeHlfV1alpha1) FabricOperatorAPIs(namespace string) v1alpha1.FabricOperatorAPIInterface {
	return &FakeFabricOperatorAPIs{c, namespace}
}

func (c *FakeHlfV1alpha1) FabricOperatorUIs(namespace string) v1alpha1.FabricOperatorUIInterface {
	return &FakeFabricOperatorUIs{c, namespace}
}

func (c *FakeHlfV1alpha1) FabricOrdererNodes(namespace string) v1alpha1.FabricOrdererNodeInterface {
	return &FakeFabricOrdererNodes{c, namespace}
}

func (c *FakeHlfV1alpha1) FabricOrderingServices(namespace string) v1alpha1.FabricOrderingServiceInterface {
	return &FakeFabricOrderingServices{c, namespace}
}

func (c *FakeHlfV1alpha1) FabricPeers(namespace string) v1alpha1.FabricPeerInterface {
	return &FakeFabricPeers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHlfV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
