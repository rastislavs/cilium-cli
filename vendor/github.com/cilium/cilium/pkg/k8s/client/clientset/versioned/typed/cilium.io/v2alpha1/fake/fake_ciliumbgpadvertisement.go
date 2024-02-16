// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumBGPAdvertisements implements CiliumBGPAdvertisementInterface
type FakeCiliumBGPAdvertisements struct {
	Fake *FakeCiliumV2alpha1
}

var ciliumbgpadvertisementsResource = v2alpha1.SchemeGroupVersion.WithResource("ciliumbgpadvertisements")

var ciliumbgpadvertisementsKind = v2alpha1.SchemeGroupVersion.WithKind("CiliumBGPAdvertisement")

// Get takes name of the ciliumBGPAdvertisement, and returns the corresponding ciliumBGPAdvertisement object, and an error if there is any.
func (c *FakeCiliumBGPAdvertisements) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumBGPAdvertisement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumbgpadvertisementsResource, name), &v2alpha1.CiliumBGPAdvertisement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPAdvertisement), err
}

// List takes label and field selectors, and returns the list of CiliumBGPAdvertisements that match those selectors.
func (c *FakeCiliumBGPAdvertisements) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumBGPAdvertisementList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumbgpadvertisementsResource, ciliumbgpadvertisementsKind, opts), &v2alpha1.CiliumBGPAdvertisementList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.CiliumBGPAdvertisementList{ListMeta: obj.(*v2alpha1.CiliumBGPAdvertisementList).ListMeta}
	for _, item := range obj.(*v2alpha1.CiliumBGPAdvertisementList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumBGPAdvertisements.
func (c *FakeCiliumBGPAdvertisements) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumbgpadvertisementsResource, opts))
}

// Create takes the representation of a ciliumBGPAdvertisement and creates it.  Returns the server's representation of the ciliumBGPAdvertisement, and an error, if there is any.
func (c *FakeCiliumBGPAdvertisements) Create(ctx context.Context, ciliumBGPAdvertisement *v2alpha1.CiliumBGPAdvertisement, opts v1.CreateOptions) (result *v2alpha1.CiliumBGPAdvertisement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumbgpadvertisementsResource, ciliumBGPAdvertisement), &v2alpha1.CiliumBGPAdvertisement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPAdvertisement), err
}

// Update takes the representation of a ciliumBGPAdvertisement and updates it. Returns the server's representation of the ciliumBGPAdvertisement, and an error, if there is any.
func (c *FakeCiliumBGPAdvertisements) Update(ctx context.Context, ciliumBGPAdvertisement *v2alpha1.CiliumBGPAdvertisement, opts v1.UpdateOptions) (result *v2alpha1.CiliumBGPAdvertisement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumbgpadvertisementsResource, ciliumBGPAdvertisement), &v2alpha1.CiliumBGPAdvertisement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPAdvertisement), err
}

// Delete takes name of the ciliumBGPAdvertisement and deletes it. Returns an error if one occurs.
func (c *FakeCiliumBGPAdvertisements) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ciliumbgpadvertisementsResource, name, opts), &v2alpha1.CiliumBGPAdvertisement{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumBGPAdvertisements) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumbgpadvertisementsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.CiliumBGPAdvertisementList{})
	return err
}

// Patch applies the patch and returns the patched ciliumBGPAdvertisement.
func (c *FakeCiliumBGPAdvertisements) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumBGPAdvertisement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumbgpadvertisementsResource, name, pt, data, subresources...), &v2alpha1.CiliumBGPAdvertisement{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPAdvertisement), err
}
