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

package v1alpha1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	apisv1alpha1 "sigs.k8s.io/kwok/pkg/apis/v1alpha1"
	scheme "sigs.k8s.io/kwok/pkg/client/clientset/versioned/scheme"
)

// MetricsGetter has a method to return a MetricInterface.
// A group's client should implement this interface.
type MetricsGetter interface {
	Metrics() MetricInterface
}

// MetricInterface has methods to work with Metric resources.
type MetricInterface interface {
	Create(ctx context.Context, metric *apisv1alpha1.Metric, opts v1.CreateOptions) (*apisv1alpha1.Metric, error)
	Update(ctx context.Context, metric *apisv1alpha1.Metric, opts v1.UpdateOptions) (*apisv1alpha1.Metric, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, metric *apisv1alpha1.Metric, opts v1.UpdateOptions) (*apisv1alpha1.Metric, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*apisv1alpha1.Metric, error)
	List(ctx context.Context, opts v1.ListOptions) (*apisv1alpha1.MetricList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apisv1alpha1.Metric, err error)
	MetricExpansion
}

// metrics implements MetricInterface
type metrics struct {
	*gentype.ClientWithList[*apisv1alpha1.Metric, *apisv1alpha1.MetricList]
}

// newMetrics returns a Metrics
func newMetrics(c *KwokV1alpha1Client) *metrics {
	return &metrics{
		gentype.NewClientWithList[*apisv1alpha1.Metric, *apisv1alpha1.MetricList](
			"metrics",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *apisv1alpha1.Metric { return &apisv1alpha1.Metric{} },
			func() *apisv1alpha1.MetricList { return &apisv1alpha1.MetricList{} },
		),
	}
}
