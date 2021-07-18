// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v3

import (
	"context"
	time "time"

	projectcalicov3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	clientset "github.com/projectcalico/api/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/projectcalico/api/pkg/client/informers_generated/externalversions/internalinterfaces"
	v3 "github.com/projectcalico/api/pkg/client/listers_generated/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StagedNetworkPolicyInformer provides access to a shared informer and lister for
// StagedNetworkPolicies.
type StagedNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v3.StagedNetworkPolicyLister
}

type stagedNetworkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStagedNetworkPolicyInformer constructs a new informer for StagedNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStagedNetworkPolicyInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStagedNetworkPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStagedNetworkPolicyInformer constructs a new informer for StagedNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStagedNetworkPolicyInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().StagedNetworkPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().StagedNetworkPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&projectcalicov3.StagedNetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *stagedNetworkPolicyInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStagedNetworkPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *stagedNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&projectcalicov3.StagedNetworkPolicy{}, f.defaultInformer)
}

func (f *stagedNetworkPolicyInformer) Lister() v3.StagedNetworkPolicyLister {
	return v3.NewStagedNetworkPolicyLister(f.Informer().GetIndexer())
}