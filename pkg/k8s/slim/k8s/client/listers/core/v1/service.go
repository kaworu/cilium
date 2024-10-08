// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ServiceLister helps list Services.
// All objects returned here must be treated as read-only.
type ServiceLister interface {
	// List lists all Services in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Service, err error)
	// Services returns an object that can list and get Services.
	Services(namespace string) ServiceNamespaceLister
	ServiceListerExpansion
}

// serviceLister implements the ServiceLister interface.
type serviceLister struct {
	listers.ResourceIndexer[*v1.Service]
}

// NewServiceLister returns a new ServiceLister.
func NewServiceLister(indexer cache.Indexer) ServiceLister {
	return &serviceLister{listers.New[*v1.Service](indexer, v1.Resource("service"))}
}

// Services returns an object that can list and get Services.
func (s *serviceLister) Services(namespace string) ServiceNamespaceLister {
	return serviceNamespaceLister{listers.NewNamespaced[*v1.Service](s.ResourceIndexer, namespace)}
}

// ServiceNamespaceLister helps list and get Services.
// All objects returned here must be treated as read-only.
type ServiceNamespaceLister interface {
	// List lists all Services in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Service, err error)
	// Get retrieves the Service from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Service, error)
	ServiceNamespaceListerExpansion
}

// serviceNamespaceLister implements the ServiceNamespaceLister
// interface.
type serviceNamespaceLister struct {
	listers.ResourceIndexer[*v1.Service]
}
