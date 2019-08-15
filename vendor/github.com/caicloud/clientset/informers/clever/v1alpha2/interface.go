/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Flavors returns a FlavorInformer.
	Flavors() FlavorInformer
	// MLNeurons returns a MLNeuronInformer.
	MLNeurons() MLNeuronInformer
	// MLNeuronTaskOwners returns a MLNeuronTaskOwnerInformer.
	MLNeuronTaskOwners() MLNeuronTaskOwnerInformer
	// Projects returns a ProjectInformer.
	Projects() ProjectInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Flavors returns a FlavorInformer.
func (v *version) Flavors() FlavorInformer {
	return &flavorInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MLNeurons returns a MLNeuronInformer.
func (v *version) MLNeurons() MLNeuronInformer {
	return &mLNeuronInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MLNeuronTaskOwners returns a MLNeuronTaskOwnerInformer.
func (v *version) MLNeuronTaskOwners() MLNeuronTaskOwnerInformer {
	return &mLNeuronTaskOwnerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Projects returns a ProjectInformer.
func (v *version) Projects() ProjectInformer {
	return &projectInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
