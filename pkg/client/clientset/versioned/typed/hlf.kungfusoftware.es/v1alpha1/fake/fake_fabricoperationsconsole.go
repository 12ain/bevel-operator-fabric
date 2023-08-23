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

// FakeFabricOperationsConsoles implements FabricOperationsConsoleInterface
type FakeFabricOperationsConsoles struct {
	Fake *FakeHlfV1alpha1
	ns   string
}

var fabricoperationsconsolesResource = v1alpha1.SchemeGroupVersion.WithResource("fabricoperationsconsoles")

var fabricoperationsconsolesKind = v1alpha1.SchemeGroupVersion.WithKind("FabricOperationsConsole")

// Get takes name of the fabricOperationsConsole, and returns the corresponding fabricOperationsConsole object, and an error if there is any.
func (c *FakeFabricOperationsConsoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricOperationsConsole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fabricoperationsconsolesResource, c.ns, name), &v1alpha1.FabricOperationsConsole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperationsConsole), err
}

// List takes label and field selectors, and returns the list of FabricOperationsConsoles that match those selectors.
func (c *FakeFabricOperationsConsoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricOperationsConsoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fabricoperationsconsolesResource, fabricoperationsconsolesKind, c.ns, opts), &v1alpha1.FabricOperationsConsoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FabricOperationsConsoleList{ListMeta: obj.(*v1alpha1.FabricOperationsConsoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.FabricOperationsConsoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fabricOperationsConsoles.
func (c *FakeFabricOperationsConsoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fabricoperationsconsolesResource, c.ns, opts))

}

// Create takes the representation of a fabricOperationsConsole and creates it.  Returns the server's representation of the fabricOperationsConsole, and an error, if there is any.
func (c *FakeFabricOperationsConsoles) Create(ctx context.Context, fabricOperationsConsole *v1alpha1.FabricOperationsConsole, opts v1.CreateOptions) (result *v1alpha1.FabricOperationsConsole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fabricoperationsconsolesResource, c.ns, fabricOperationsConsole), &v1alpha1.FabricOperationsConsole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperationsConsole), err
}

// Update takes the representation of a fabricOperationsConsole and updates it. Returns the server's representation of the fabricOperationsConsole, and an error, if there is any.
func (c *FakeFabricOperationsConsoles) Update(ctx context.Context, fabricOperationsConsole *v1alpha1.FabricOperationsConsole, opts v1.UpdateOptions) (result *v1alpha1.FabricOperationsConsole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fabricoperationsconsolesResource, c.ns, fabricOperationsConsole), &v1alpha1.FabricOperationsConsole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperationsConsole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFabricOperationsConsoles) UpdateStatus(ctx context.Context, fabricOperationsConsole *v1alpha1.FabricOperationsConsole, opts v1.UpdateOptions) (*v1alpha1.FabricOperationsConsole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fabricoperationsconsolesResource, "status", c.ns, fabricOperationsConsole), &v1alpha1.FabricOperationsConsole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperationsConsole), err
}

// Delete takes name of the fabricOperationsConsole and deletes it. Returns an error if one occurs.
func (c *FakeFabricOperationsConsoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(fabricoperationsconsolesResource, c.ns, name, opts), &v1alpha1.FabricOperationsConsole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFabricOperationsConsoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fabricoperationsconsolesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FabricOperationsConsoleList{})
	return err
}

// Patch applies the patch and returns the patched fabricOperationsConsole.
func (c *FakeFabricOperationsConsoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricOperationsConsole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricoperationsconsolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FabricOperationsConsole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperationsConsole), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fabricOperationsConsole.
func (c *FakeFabricOperationsConsoles) Apply(ctx context.Context, fabricOperationsConsole *hlfkungfusoftwareesv1alpha1.FabricOperationsConsoleApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricOperationsConsole, err error) {
	if fabricOperationsConsole == nil {
		return nil, fmt.Errorf("fabricOperationsConsole provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricOperationsConsole)
	if err != nil {
		return nil, err
	}
	name := fabricOperationsConsole.Name
	if name == nil {
		return nil, fmt.Errorf("fabricOperationsConsole.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricoperationsconsolesResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.FabricOperationsConsole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperationsConsole), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeFabricOperationsConsoles) ApplyStatus(ctx context.Context, fabricOperationsConsole *hlfkungfusoftwareesv1alpha1.FabricOperationsConsoleApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricOperationsConsole, err error) {
	if fabricOperationsConsole == nil {
		return nil, fmt.Errorf("fabricOperationsConsole provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricOperationsConsole)
	if err != nil {
		return nil, err
	}
	name := fabricOperationsConsole.Name
	if name == nil {
		return nil, fmt.Errorf("fabricOperationsConsole.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricoperationsconsolesResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.FabricOperationsConsole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperationsConsole), err
}
