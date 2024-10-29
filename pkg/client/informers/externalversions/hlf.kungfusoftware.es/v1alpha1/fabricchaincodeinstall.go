/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	hlfkungfusoftwareesv1alpha1 "github.com/kfsoftware/hlf-operator/pkg/apis/hlf.kungfusoftware.es/v1alpha1"
	versioned "github.com/kfsoftware/hlf-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kfsoftware/hlf-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kfsoftware/hlf-operator/pkg/client/listers/hlf.kungfusoftware.es/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FabricChaincodeInstallInformer provides access to a shared informer and lister for
// FabricChaincodeInstalls.
type FabricChaincodeInstallInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FabricChaincodeInstallLister
}

type fabricChaincodeInstallInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewFabricChaincodeInstallInformer constructs a new informer for FabricChaincodeInstall type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFabricChaincodeInstallInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFabricChaincodeInstallInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredFabricChaincodeInstallInformer constructs a new informer for FabricChaincodeInstall type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFabricChaincodeInstallInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HlfV1alpha1().FabricChaincodeInstalls().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HlfV1alpha1().FabricChaincodeInstalls().Watch(context.TODO(), options)
			},
		},
		&hlfkungfusoftwareesv1alpha1.FabricChaincodeInstall{},
		resyncPeriod,
		indexers,
	)
}

func (f *fabricChaincodeInstallInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFabricChaincodeInstallInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fabricChaincodeInstallInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hlfkungfusoftwareesv1alpha1.FabricChaincodeInstall{}, f.defaultInformer)
}

func (f *fabricChaincodeInstallInformer) Lister() v1alpha1.FabricChaincodeInstallLister {
	return v1alpha1.NewFabricChaincodeInstallLister(f.Informer().GetIndexer())
}
