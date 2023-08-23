/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	hlfkungfusoftwareesv1alpha1 "github.com/kfsoftware/hlf-operator/pkg/client/applyconfiguration/hlf.kungfusoftware.es/v1alpha1"
	scheme "github.com/kfsoftware/hlf-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FabricMainChannelsGetter has a method to return a FabricMainChannelInterface.
// A group's client should implement this interface.
type FabricMainChannelsGetter interface {
	FabricMainChannels() FabricMainChannelInterface
}

// FabricMainChannelInterface has methods to work with FabricMainChannel resources.
type FabricMainChannelInterface interface {
	Create(ctx context.Context, fabricMainChannel *v1alpha1.FabricMainChannel, opts v1.CreateOptions) (*v1alpha1.FabricMainChannel, error)
	Update(ctx context.Context, fabricMainChannel *v1alpha1.FabricMainChannel, opts v1.UpdateOptions) (*v1alpha1.FabricMainChannel, error)
	UpdateStatus(ctx context.Context, fabricMainChannel *v1alpha1.FabricMainChannel, opts v1.UpdateOptions) (*v1alpha1.FabricMainChannel, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FabricMainChannel, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FabricMainChannelList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricMainChannel, err error)
	Apply(ctx context.Context, fabricMainChannel *hlfkungfusoftwareesv1alpha1.FabricMainChannelApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricMainChannel, err error)
	ApplyStatus(ctx context.Context, fabricMainChannel *hlfkungfusoftwareesv1alpha1.FabricMainChannelApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricMainChannel, err error)
	FabricMainChannelExpansion
}

// fabricMainChannels implements FabricMainChannelInterface
type fabricMainChannels struct {
	client rest.Interface
}

// newFabricMainChannels returns a FabricMainChannels
func newFabricMainChannels(c *HlfV1alpha1Client) *fabricMainChannels {
	return &fabricMainChannels{
		client: c.RESTClient(),
	}
}

// Get takes name of the fabricMainChannel, and returns the corresponding fabricMainChannel object, and an error if there is any.
func (c *fabricMainChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricMainChannel, err error) {
	result = &v1alpha1.FabricMainChannel{}
	err = c.client.Get().
		Resource("fabricmainchannels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FabricMainChannels that match those selectors.
func (c *fabricMainChannels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricMainChannelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FabricMainChannelList{}
	err = c.client.Get().
		Resource("fabricmainchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fabricMainChannels.
func (c *fabricMainChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("fabricmainchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a fabricMainChannel and creates it.  Returns the server's representation of the fabricMainChannel, and an error, if there is any.
func (c *fabricMainChannels) Create(ctx context.Context, fabricMainChannel *v1alpha1.FabricMainChannel, opts v1.CreateOptions) (result *v1alpha1.FabricMainChannel, err error) {
	result = &v1alpha1.FabricMainChannel{}
	err = c.client.Post().
		Resource("fabricmainchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fabricMainChannel).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a fabricMainChannel and updates it. Returns the server's representation of the fabricMainChannel, and an error, if there is any.
func (c *fabricMainChannels) Update(ctx context.Context, fabricMainChannel *v1alpha1.FabricMainChannel, opts v1.UpdateOptions) (result *v1alpha1.FabricMainChannel, err error) {
	result = &v1alpha1.FabricMainChannel{}
	err = c.client.Put().
		Resource("fabricmainchannels").
		Name(fabricMainChannel.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fabricMainChannel).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *fabricMainChannels) UpdateStatus(ctx context.Context, fabricMainChannel *v1alpha1.FabricMainChannel, opts v1.UpdateOptions) (result *v1alpha1.FabricMainChannel, err error) {
	result = &v1alpha1.FabricMainChannel{}
	err = c.client.Put().
		Resource("fabricmainchannels").
		Name(fabricMainChannel.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fabricMainChannel).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the fabricMainChannel and deletes it. Returns an error if one occurs.
func (c *fabricMainChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("fabricmainchannels").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fabricMainChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("fabricmainchannels").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched fabricMainChannel.
func (c *fabricMainChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricMainChannel, err error) {
	result = &v1alpha1.FabricMainChannel{}
	err = c.client.Patch(pt).
		Resource("fabricmainchannels").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fabricMainChannel.
func (c *fabricMainChannels) Apply(ctx context.Context, fabricMainChannel *hlfkungfusoftwareesv1alpha1.FabricMainChannelApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricMainChannel, err error) {
	if fabricMainChannel == nil {
		return nil, fmt.Errorf("fabricMainChannel provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(fabricMainChannel)
	if err != nil {
		return nil, err
	}
	name := fabricMainChannel.Name
	if name == nil {
		return nil, fmt.Errorf("fabricMainChannel.Name must be provided to Apply")
	}
	result = &v1alpha1.FabricMainChannel{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("fabricmainchannels").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *fabricMainChannels) ApplyStatus(ctx context.Context, fabricMainChannel *hlfkungfusoftwareesv1alpha1.FabricMainChannelApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricMainChannel, err error) {
	if fabricMainChannel == nil {
		return nil, fmt.Errorf("fabricMainChannel provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(fabricMainChannel)
	if err != nil {
		return nil, err
	}

	name := fabricMainChannel.Name
	if name == nil {
		return nil, fmt.Errorf("fabricMainChannel.Name must be provided to Apply")
	}

	result = &v1alpha1.FabricMainChannel{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("fabricmainchannels").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
