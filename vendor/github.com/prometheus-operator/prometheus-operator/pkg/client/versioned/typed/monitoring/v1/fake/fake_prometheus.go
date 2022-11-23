// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrometheuses implements PrometheusInterface
type FakePrometheuses struct {
	Fake *FakeMonitoringV1
	ns   string
}

var prometheusesResource = schema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1", Resource: "prometheuses"}

var prometheusesKind = schema.GroupVersionKind{Group: "monitoring.coreos.com", Version: "v1", Kind: "Prometheus"}

// Get takes name of the prometheus, and returns the corresponding prometheus object, and an error if there is any.
func (c *FakePrometheuses) Get(ctx context.Context, name string, options v1.GetOptions) (result *monitoringv1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(prometheusesResource, c.ns, name), &monitoringv1.Prometheus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.Prometheus), err
}

// List takes label and field selectors, and returns the list of Prometheuses that match those selectors.
func (c *FakePrometheuses) List(ctx context.Context, opts v1.ListOptions) (result *monitoringv1.PrometheusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(prometheusesResource, prometheusesKind, c.ns, opts), &monitoringv1.PrometheusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoringv1.PrometheusList{ListMeta: obj.(*monitoringv1.PrometheusList).ListMeta}
	for _, item := range obj.(*monitoringv1.PrometheusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested prometheuses.
func (c *FakePrometheuses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(prometheusesResource, c.ns, opts))

}

// Create takes the representation of a prometheus and creates it.  Returns the server's representation of the prometheus, and an error, if there is any.
func (c *FakePrometheuses) Create(ctx context.Context, prometheus *monitoringv1.Prometheus, opts v1.CreateOptions) (result *monitoringv1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(prometheusesResource, c.ns, prometheus), &monitoringv1.Prometheus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.Prometheus), err
}

// Update takes the representation of a prometheus and updates it. Returns the server's representation of the prometheus, and an error, if there is any.
func (c *FakePrometheuses) Update(ctx context.Context, prometheus *monitoringv1.Prometheus, opts v1.UpdateOptions) (result *monitoringv1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(prometheusesResource, c.ns, prometheus), &monitoringv1.Prometheus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.Prometheus), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrometheuses) UpdateStatus(ctx context.Context, prometheus *monitoringv1.Prometheus, opts v1.UpdateOptions) (*monitoringv1.Prometheus, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(prometheusesResource, "status", c.ns, prometheus), &monitoringv1.Prometheus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.Prometheus), err
}

// Delete takes name of the prometheus and deletes it. Returns an error if one occurs.
func (c *FakePrometheuses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(prometheusesResource, c.ns, name, opts), &monitoringv1.Prometheus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrometheuses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(prometheusesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &monitoringv1.PrometheusList{})
	return err
}

// Patch applies the patch and returns the patched prometheus.
func (c *FakePrometheuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *monitoringv1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(prometheusesResource, c.ns, name, pt, data, subresources...), &monitoringv1.Prometheus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.Prometheus), err
}
