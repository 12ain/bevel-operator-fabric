/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFabricOperatorAPIs implements FabricOperatorAPIInterface
type FakeFabricOperatorAPIs struct {
	Fake *FakeHlfV1alpha1
	ns   string
}

var fabricoperatorapisResource = schema.GroupVersionResource{Group: "hlf.kungfusoftware.es", Version: "v1alpha1", Resource: "fabricoperatorapis"}

var fabricoperatorapisKind = schema.GroupVersionKind{Group: "hlf.kungfusoftware.es", Version: "v1alpha1", Kind: "FabricOperatorAPI"}

// Get takes name of the fabricOperatorAPI, and returns the corresponding fabricOperatorAPI object, and an error if there is any.
func (c *FakeFabricOperatorAPIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricOperatorAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fabricoperatorapisResource, c.ns, name), &v1alpha1.FabricOperatorAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorAPI), err
}

// List takes label and field selectors, and returns the list of FabricOperatorAPIs that match those selectors.
func (c *FakeFabricOperatorAPIs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricOperatorAPIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fabricoperatorapisResource, fabricoperatorapisKind, c.ns, opts), &v1alpha1.FabricOperatorAPIList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FabricOperatorAPIList{ListMeta: obj.(*v1alpha1.FabricOperatorAPIList).ListMeta}
	for _, item := range obj.(*v1alpha1.FabricOperatorAPIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fabricOperatorAPIs.
func (c *FakeFabricOperatorAPIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fabricoperatorapisResource, c.ns, opts))

}

// Create takes the representation of a fabricOperatorAPI and creates it.  Returns the server's representation of the fabricOperatorAPI, and an error, if there is any.
func (c *FakeFabricOperatorAPIs) Create(ctx context.Context, fabricOperatorAPI *v1alpha1.FabricOperatorAPI, opts v1.CreateOptions) (result *v1alpha1.FabricOperatorAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fabricoperatorapisResource, c.ns, fabricOperatorAPI), &v1alpha1.FabricOperatorAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorAPI), err
}

// Update takes the representation of a fabricOperatorAPI and updates it. Returns the server's representation of the fabricOperatorAPI, and an error, if there is any.
func (c *FakeFabricOperatorAPIs) Update(ctx context.Context, fabricOperatorAPI *v1alpha1.FabricOperatorAPI, opts v1.UpdateOptions) (result *v1alpha1.FabricOperatorAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fabricoperatorapisResource, c.ns, fabricOperatorAPI), &v1alpha1.FabricOperatorAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorAPI), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFabricOperatorAPIs) UpdateStatus(ctx context.Context, fabricOperatorAPI *v1alpha1.FabricOperatorAPI, opts v1.UpdateOptions) (*v1alpha1.FabricOperatorAPI, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fabricoperatorapisResource, "status", c.ns, fabricOperatorAPI), &v1alpha1.FabricOperatorAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorAPI), err
}

// Delete takes name of the fabricOperatorAPI and deletes it. Returns an error if one occurs.
func (c *FakeFabricOperatorAPIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fabricoperatorapisResource, c.ns, name), &v1alpha1.FabricOperatorAPI{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFabricOperatorAPIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fabricoperatorapisResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FabricOperatorAPIList{})
	return err
}

// Patch applies the patch and returns the patched fabricOperatorAPI.
func (c *FakeFabricOperatorAPIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricOperatorAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricoperatorapisResource, c.ns, name, pt, data, subresources...), &v1alpha1.FabricOperatorAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorAPI), err
}