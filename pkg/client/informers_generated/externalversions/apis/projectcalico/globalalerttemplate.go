// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package projectcalico

import (
	time "time"

	apisprojectcalico "github.com/tigera/api/pkg/apis/projectcalico"
	clientset "github.com/tigera/api/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/tigera/api/pkg/client/informers_generated/externalversions/internalinterfaces"
	projectcalico "github.com/tigera/api/pkg/client/listers_generated/apis/projectcalico"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GlobalAlertTemplateInformer provides access to a shared informer and lister for
// GlobalAlertTemplates.
type GlobalAlertTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() projectcalico.GlobalAlertTemplateLister
}

type globalAlertTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewGlobalAlertTemplateInformer constructs a new informer for GlobalAlertTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGlobalAlertTemplateInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGlobalAlertTemplateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredGlobalAlertTemplateInformer constructs a new informer for GlobalAlertTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGlobalAlertTemplateInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoProjectcalico().GlobalAlertTemplates().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoProjectcalico().GlobalAlertTemplates().Watch(options)
			},
		},
		&apisprojectcalico.GlobalAlertTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *globalAlertTemplateInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGlobalAlertTemplateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *globalAlertTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisprojectcalico.GlobalAlertTemplate{}, f.defaultInformer)
}

func (f *globalAlertTemplateInformer) Lister() projectcalico.GlobalAlertTemplateLister {
	return projectcalico.NewGlobalAlertTemplateLister(f.Informer().GetIndexer())
}