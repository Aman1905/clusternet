/*
Copyright The Clusternet Authors.

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
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	appsv1alpha1 "github.com/clusternet/clusternet/pkg/generated/applyconfiguration/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSubscriptions implements SubscriptionInterface
type FakeSubscriptions struct {
	Fake *FakeAppsV1alpha1
	ns   string
}

var subscriptionsResource = v1alpha1.SchemeGroupVersion.WithResource("subscriptions")

var subscriptionsKind = v1alpha1.SchemeGroupVersion.WithKind("Subscription")

// Get takes name of the subscription, and returns the corresponding subscription object, and an error if there is any.
func (c *FakeSubscriptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Subscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(subscriptionsResource, c.ns, name), &v1alpha1.Subscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subscription), err
}

// List takes label and field selectors, and returns the list of Subscriptions that match those selectors.
func (c *FakeSubscriptions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(subscriptionsResource, subscriptionsKind, c.ns, opts), &v1alpha1.SubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubscriptionList{ListMeta: obj.(*v1alpha1.SubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested subscriptions.
func (c *FakeSubscriptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(subscriptionsResource, c.ns, opts))

}

// Create takes the representation of a subscription and creates it.  Returns the server's representation of the subscription, and an error, if there is any.
func (c *FakeSubscriptions) Create(ctx context.Context, subscription *v1alpha1.Subscription, opts v1.CreateOptions) (result *v1alpha1.Subscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(subscriptionsResource, c.ns, subscription), &v1alpha1.Subscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subscription), err
}

// Update takes the representation of a subscription and updates it. Returns the server's representation of the subscription, and an error, if there is any.
func (c *FakeSubscriptions) Update(ctx context.Context, subscription *v1alpha1.Subscription, opts v1.UpdateOptions) (result *v1alpha1.Subscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(subscriptionsResource, c.ns, subscription), &v1alpha1.Subscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubscriptions) UpdateStatus(ctx context.Context, subscription *v1alpha1.Subscription, opts v1.UpdateOptions) (*v1alpha1.Subscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(subscriptionsResource, "status", c.ns, subscription), &v1alpha1.Subscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subscription), err
}

// Delete takes name of the subscription and deletes it. Returns an error if one occurs.
func (c *FakeSubscriptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(subscriptionsResource, c.ns, name, opts), &v1alpha1.Subscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubscriptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(subscriptionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched subscription.
func (c *FakeSubscriptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Subscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Subscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subscription), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied subscription.
func (c *FakeSubscriptions) Apply(ctx context.Context, subscription *appsv1alpha1.SubscriptionApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Subscription, err error) {
	if subscription == nil {
		return nil, fmt.Errorf("subscription provided to Apply must not be nil")
	}
	data, err := json.Marshal(subscription)
	if err != nil {
		return nil, err
	}
	name := subscription.Name
	if name == nil {
		return nil, fmt.Errorf("subscription.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subscriptionsResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.Subscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subscription), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeSubscriptions) ApplyStatus(ctx context.Context, subscription *appsv1alpha1.SubscriptionApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Subscription, err error) {
	if subscription == nil {
		return nil, fmt.Errorf("subscription provided to Apply must not be nil")
	}
	data, err := json.Marshal(subscription)
	if err != nil {
		return nil, err
	}
	name := subscription.Name
	if name == nil {
		return nil, fmt.Errorf("subscription.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subscriptionsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.Subscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Subscription), err
}
