/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v2beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	informers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
	v2beta1 "k8s.io/client-go/listers/autoscaling/v2beta1"
)

// Interface provides access to all the listers in this group version.
type Interface interface { // HorizontalPodAutoscalers returns a HorizontalPodAutoscalerLister
	HorizontalPodAutoscalers() v2beta1.HorizontalPodAutoscalerLister
}

type version struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

type infromerVersion struct {
	factory informers.SharedInformerFactory
}

// New returns a new Interface.
func New(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{client: client, tweakListOptions: tweakListOptions}
}

// NewFrom returns a new Interface.
func NewFrom(factory informers.SharedInformerFactory) Interface {
	return &infromerVersion{factory: factory}
}

// HorizontalPodAutoscalers returns a HorizontalPodAutoscalerLister.
func (v *version) HorizontalPodAutoscalers() v2beta1.HorizontalPodAutoscalerLister {
	return &horizontalPodAutoscalerLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// HorizontalPodAutoscalers returns a HorizontalPodAutoscalerLister.
func (v *infromerVersion) HorizontalPodAutoscalers() v2beta1.HorizontalPodAutoscalerLister {
	return v.factory.Autoscaling().V2beta1().HorizontalPodAutoscalers().Lister()
}
