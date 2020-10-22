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

package ingress

import (
	"k8s.io/apimachinery/pkg/util/intstr"
	"knative.dev/networking/pkg/apis/networking"
	"knative.dev/networking/pkg/apis/networking/v1alpha1"
	"knative.dev/networking/test"
)

// TestRewriteHost verifies that a RewriteHost rule can be used to implement vanity URLs.
func TestRewriteHost(t *test.T) {
	t.Parallel()

	name, port, _ := CreateRuntimeService(t.C, t, t.Clients, networking.ServicePortNameHTTP1)

	privateServiceName := test.ObjectNameForTest(t)
	privateHostName := privateServiceName + "." + t.TestNamespace + ".svc.cluster.local"

	// Create a simple Ingress over the Service.
	ing, _, _ := CreateIngressReady(t.C, t, t.Clients, v1alpha1.IngressSpec{
		Rules: []v1alpha1.IngressRule{{
			Visibility: v1alpha1.IngressVisibilityClusterLocal,
			Hosts:      []string{privateHostName},
			HTTP: &v1alpha1.HTTPIngressRuleValue{
				Paths: []v1alpha1.HTTPIngressPath{{
					Splits: []v1alpha1.IngressBackendSplit{{
						IngressBackend: v1alpha1.IngressBackend{
							ServiceName:      name,
							ServiceNamespace: t.TestNamespace,
							ServicePort:      intstr.FromInt(port),
						},
					}},
				}},
			},
		}},
	})

	// Slap an ExternalName service in front of the kingress
	loadbalancerAddress := ing.Status.PrivateLoadBalancer.Ingress[0].DomainInternal
	createExternalNameService(t.C, t, t.Clients, privateHostName, loadbalancerAddress)

	hosts := []string{
		"vanity.ismy.name",
		"vanity.isalsomy.number",
	}

	// Using fixed hostnames can lead to conflicts when -count=N>1
	// so pseudo-randomize the hostnames to avoid conflicts.
	for i, host := range hosts {
		hosts[i] = name + "." + host
	}

	// Now create a RewriteHost ingress to point a custom Host at the Service
	_, client, _ := CreateIngressReady(t.C, t, t.Clients, v1alpha1.IngressSpec{
		Rules: []v1alpha1.IngressRule{{
			Hosts:      hosts,
			Visibility: v1alpha1.IngressVisibilityExternalIP,
			HTTP: &v1alpha1.HTTPIngressRuleValue{
				Paths: []v1alpha1.HTTPIngressPath{{
					RewriteHost: privateHostName,
					Splits: []v1alpha1.IngressBackendSplit{{
						IngressBackend: v1alpha1.IngressBackend{
							ServiceName:      privateServiceName,
							ServiceNamespace: t.TestNamespace,
							ServicePort:      intstr.FromInt(80),
						},
					}},
				}},
			},
		}},
	})

	for _, host := range hosts {
		RuntimeRequest(t.C, t, client, "http://"+host)
	}
}
