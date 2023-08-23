/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	hlfkungfusoftwareesv1alpha1 "github.com/kfsoftware/hlf-operator/pkg/client/applyconfiguration/hlf.kungfusoftware.es/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFabricCAs implements FabricCAInterface
type FakeFabricCAs struct {
	Fake *FakeHlfV1alpha1
	ns   string
}

var fabriccasResource = v1alpha1.SchemeGroupVersion.WithResource("fabriccas")

var fabriccasKind = v1alpha1.SchemeGroupVersion.WithKind("FabricCA")

// Get takes name of the fabricCA, and returns the corresponding fabricCA object, and an error if there is any.
func (c *FakeFabricCAs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fabriccasResource, c.ns, name), &v1alpha1.FabricCA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricCA), err
}

// List takes label and field selectors, and returns the list of FabricCAs that match those selectors.
func (c *FakeFabricCAs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricCAList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fabriccasResource, fabriccasKind, c.ns, opts), &v1alpha1.FabricCAList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FabricCAList{ListMeta: obj.(*v1alpha1.FabricCAList).ListMeta}
	for _, item := range obj.(*v1alpha1.FabricCAList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fabricCAs.
func (c *FakeFabricCAs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fabriccasResource, c.ns, opts))

}

// Create takes the representation of a fabricCA and creates it.  Returns the server's representation of the fabricCA, and an error, if there is any.
func (c *FakeFabricCAs) Create(ctx context.Context, fabricCA *v1alpha1.FabricCA, opts v1.CreateOptions) (result *v1alpha1.FabricCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fabriccasResource, c.ns, fabricCA), &v1alpha1.FabricCA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricCA), err
}

// Update takes the representation of a fabricCA and updates it. Returns the server's representation of the fabricCA, and an error, if there is any.
func (c *FakeFabricCAs) Update(ctx context.Context, fabricCA *v1alpha1.FabricCA, opts v1.UpdateOptions) (result *v1alpha1.FabricCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fabriccasResource, c.ns, fabricCA), &v1alpha1.FabricCA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricCA), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFabricCAs) UpdateStatus(ctx context.Context, fabricCA *v1alpha1.FabricCA, opts v1.UpdateOptions) (*v1alpha1.FabricCA, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fabriccasResource, "status", c.ns, fabricCA), &v1alpha1.FabricCA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricCA), err
}

// Delete takes name of the fabricCA and deletes it. Returns an error if one occurs.
func (c *FakeFabricCAs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(fabriccasResource, c.ns, name, opts), &v1alpha1.FabricCA{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFabricCAs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fabriccasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FabricCAList{})
	return err
}

// Patch applies the patch and returns the patched fabricCA.
func (c *FakeFabricCAs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabriccasResource, c.ns, name, pt, data, subresources...), &v1alpha1.FabricCA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricCA), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fabricCA.
func (c *FakeFabricCAs) Apply(ctx context.Context, fabricCA *hlfkungfusoftwareesv1alpha1.FabricCAApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricCA, err error) {
	if fabricCA == nil {
		return nil, fmt.Errorf("fabricCA provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricCA)
	if err != nil {
		return nil, err
	}
	name := fabricCA.Name
	if name == nil {
		return nil, fmt.Errorf("fabricCA.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabriccasResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.FabricCA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricCA), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeFabricCAs) ApplyStatus(ctx context.Context, fabricCA *hlfkungfusoftwareesv1alpha1.FabricCAApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricCA, err error) {
	if fabricCA == nil {
		return nil, fmt.Errorf("fabricCA provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricCA)
	if err != nil {
		return nil, err
	}
	name := fabricCA.Name
	if name == nil {
		return nil, fmt.Errorf("fabricCA.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabriccasResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.FabricCA{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricCA), err
}
