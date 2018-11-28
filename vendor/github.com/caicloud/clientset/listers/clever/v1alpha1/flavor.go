/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/clever/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FlavorLister helps list Flavors.
type FlavorLister interface {
	// List lists all Flavors in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Flavor, err error)
	// Get retrieves the Flavor from the index for a given name.
	Get(name string) (*v1alpha1.Flavor, error)
	FlavorListerExpansion
}

// flavorLister implements the FlavorLister interface.
type flavorLister struct {
	indexer cache.Indexer
}

// NewFlavorLister returns a new FlavorLister.
func NewFlavorLister(indexer cache.Indexer) FlavorLister {
	return &flavorLister{indexer: indexer}
}

// List lists all Flavors in the indexer.
func (s *flavorLister) List(selector labels.Selector) (ret []*v1alpha1.Flavor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Flavor))
	})
	return ret, err
}

// Get retrieves the Flavor from the index for a given name.
func (s *flavorLister) Get(name string) (*v1alpha1.Flavor, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("flavor"), name)
	}
	return obj.(*v1alpha1.Flavor), nil
}
