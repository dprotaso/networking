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

package conformance

import (
	"context"
	networkingtest "knative.dev/networking/test"
)

type tKey struct{}

func WithT(ctx context.Context, t *networkingtest.T) context.Context {
	return context.WithValue(ctx, tKey{}, t)
}

func TFromContext(ctx context.Context) *networkingtest.T {
	if t, ok := ctx.Value(tKey{}).(*networkingtest.T); ok {
		return t
	}
	panic("no T found in context")
}
