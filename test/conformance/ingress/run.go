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
	rigging "knative.dev/reconciler-test/rigging/v2"
)

// RunConformance will run ingress conformance tests
//
// Depending on the options it may test alpha and beta features
func Conformance() *rigging.Feature {
	f := new(rigging.Feature)

	f.Stable("Ingress Conformance").
		Must("basics", TestBasics).
		Must("basics/http2", TestBasicsHTTP2).
		Must("grpc", TestGRPC).
		Must("grpc/split", TestGRPCSplit).
		Must("headers/pre-split", TestPreSplitSetHeaders).
		Must("headers/post-split", TestPostSplitSetHeaders).
		Must("hosts/multiple", TestMultipleHosts).
		Must("dispatch/path", TestPath).
		Must("dispatch/percentage", TestPercentage).
		Must("dispatch/path_and_percentage", TestPathAndPercentageSplit).
		Must("timeout", TestTimeout).
		Must("tls", TestIngressTLS).
		Must("update", TestUpdate).
		Must("visibility", TestVisibility).
		Must("visibility/split", TestVisibilitySplit).
		Must("visibility/path", TestVisibilityPath).
		Must("ingressclass", TestIngressClass).
		Must("websocket", TestWebsocket).
		Must("websocket/split", TestWebsocketSplit)

	f.Beta("Conformance").
		Should("headers/probe", TestProbeHeaders)

	f.Alpha("Conformance").
		May("headers/tag", TestTagHeaders).
		May("host-rewrite", TestRewriteHost)

	return f
}
