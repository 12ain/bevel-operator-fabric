/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/pkg/apis/hlf.kungfusoftware.es/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// FabricOperatorAPILister helps list FabricOperatorAPIs.
// All objects returned here must be treated as read-only.
type FabricOperatorAPILister interface {
	// List lists all FabricOperatorAPIs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricOperatorAPI, err error)
	// FabricOperatorAPIs returns an object that can list and get FabricOperatorAPIs.
	FabricOperatorAPIs(namespace string) FabricOperatorAPINamespaceLister
	FabricOperatorAPIListerExpansion
}

// fabricOperatorAPILister implements the FabricOperatorAPILister interface.
type fabricOperatorAPILister struct {
	listers.ResourceIndexer[*v1alpha1.FabricOperatorAPI]
}

// NewFabricOperatorAPILister returns a new FabricOperatorAPILister.
func NewFabricOperatorAPILister(indexer cache.Indexer) FabricOperatorAPILister {
	return &fabricOperatorAPILister{listers.New[*v1alpha1.FabricOperatorAPI](indexer, v1alpha1.Resource("fabricoperatorapi"))}
}

// FabricOperatorAPIs returns an object that can list and get FabricOperatorAPIs.
func (s *fabricOperatorAPILister) FabricOperatorAPIs(namespace string) FabricOperatorAPINamespaceLister {
	return fabricOperatorAPINamespaceLister{listers.NewNamespaced[*v1alpha1.FabricOperatorAPI](s.ResourceIndexer, namespace)}
}

// FabricOperatorAPINamespaceLister helps list and get FabricOperatorAPIs.
// All objects returned here must be treated as read-only.
type FabricOperatorAPINamespaceLister interface {
	// List lists all FabricOperatorAPIs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricOperatorAPI, err error)
	// Get retrieves the FabricOperatorAPI from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricOperatorAPI, error)
	FabricOperatorAPINamespaceListerExpansion
}

// fabricOperatorAPINamespaceLister implements the FabricOperatorAPINamespaceLister
// interface.
type fabricOperatorAPINamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.FabricOperatorAPI]
}
