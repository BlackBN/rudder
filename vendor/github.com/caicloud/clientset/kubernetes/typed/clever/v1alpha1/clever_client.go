/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/caicloud/clientset/kubernetes/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/clever/v1alpha1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type CleverV1alpha1Interface interface {
	RESTClient() rest.Interface
	FlavorsGetter
	ProjectsGetter
	TemplatesGetter
}

// CleverV1alpha1Client is used to interact with features provided by the clever.caicloud.io group.
type CleverV1alpha1Client struct {
	restClient rest.Interface
}

func (c *CleverV1alpha1Client) Flavors() FlavorInterface {
	return newFlavors(c)
}

func (c *CleverV1alpha1Client) Projects(namespace string) ProjectInterface {
	return newProjects(c, namespace)
}

func (c *CleverV1alpha1Client) Templates() TemplateInterface {
	return newTemplates(c)
}

// NewForConfig creates a new CleverV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*CleverV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CleverV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new CleverV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CleverV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CleverV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *CleverV1alpha1Client {
	return &CleverV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CleverV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
