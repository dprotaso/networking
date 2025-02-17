/*
Copyright 2020 The Knative Authors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	networkingv1alpha1 "knative.dev/networking/pkg/apis/networking/v1alpha1"
)

// ClusterDomainClaimLister helps list ClusterDomainClaims.
// All objects returned here must be treated as read-only.
type ClusterDomainClaimLister interface {
	// List lists all ClusterDomainClaims in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkingv1alpha1.ClusterDomainClaim, err error)
	// Get retrieves the ClusterDomainClaim from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*networkingv1alpha1.ClusterDomainClaim, error)
	ClusterDomainClaimListerExpansion
}

// clusterDomainClaimLister implements the ClusterDomainClaimLister interface.
type clusterDomainClaimLister struct {
	listers.ResourceIndexer[*networkingv1alpha1.ClusterDomainClaim]
}

// NewClusterDomainClaimLister returns a new ClusterDomainClaimLister.
func NewClusterDomainClaimLister(indexer cache.Indexer) ClusterDomainClaimLister {
	return &clusterDomainClaimLister{listers.New[*networkingv1alpha1.ClusterDomainClaim](indexer, networkingv1alpha1.Resource("clusterdomainclaim"))}
}
