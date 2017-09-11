/*
Copyright 2017 caicloud authors. All rights reserved.
*/

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ConfigClaims returns a ConfigClaimInformer.
	ConfigClaims() ConfigClaimInformer
	// ConfigReferences returns a ConfigReferenceInformer.
	ConfigReferences() ConfigReferenceInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// ConfigClaims returns a ConfigClaimInformer.
func (v *version) ConfigClaims() ConfigClaimInformer {
	return &configClaimInformer{factory: v.SharedInformerFactory}
}

// ConfigReferences returns a ConfigReferenceInformer.
func (v *version) ConfigReferences() ConfigReferenceInformer {
	return &configReferenceInformer{factory: v.SharedInformerFactory}
}
