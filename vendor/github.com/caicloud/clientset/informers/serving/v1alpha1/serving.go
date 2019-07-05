/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1alpha1 "github.com/caicloud/clientset/listers/serving/v1alpha1"
	servingv1alpha1 "github.com/caicloud/clientset/pkg/apis/serving/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// ServingInformer provides access to a shared informer and lister for
// Servings.
type ServingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServingLister
}

type servingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServingInformer constructs a new informer for Serving type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServingInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServingInformer constructs a new informer for Serving type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServingInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1alpha1().Servings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1alpha1().Servings(namespace).Watch(options)
			},
		},
		&servingv1alpha1.Serving{},
		resyncPeriod,
		indexers,
	)
}

func (f *servingInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServingInformer(client.(kubernetes.Interface), f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *servingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servingv1alpha1.Serving{}, f.defaultInformer)
}

func (f *servingInformer) Lister() v1alpha1.ServingLister {
	return v1alpha1.NewServingLister(f.Informer().GetIndexer())
}