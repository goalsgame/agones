// Copyright 2022 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "agones.dev/agones/pkg/apis/allocation/v1"
	scheme "agones.dev/agones/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// GameServerAllocationsGetter has a method to return a GameServerAllocationInterface.
// A group's client should implement this interface.
type GameServerAllocationsGetter interface {
	GameServerAllocations(namespace string) GameServerAllocationInterface
}

// GameServerAllocationInterface has methods to work with GameServerAllocation resources.
type GameServerAllocationInterface interface {
	Create(ctx context.Context, gameServerAllocation *v1.GameServerAllocation, opts metav1.CreateOptions) (*v1.GameServerAllocation, error)
	GameServerAllocationExpansion
}

// gameServerAllocations implements GameServerAllocationInterface
type gameServerAllocations struct {
	client rest.Interface
	ns     string
}

// newGameServerAllocations returns a GameServerAllocations
func newGameServerAllocations(c *AllocationV1Client, namespace string) *gameServerAllocations {
	return &gameServerAllocations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a gameServerAllocation and creates it.  Returns the server's representation of the gameServerAllocation, and an error, if there is any.
func (c *gameServerAllocations) Create(ctx context.Context, gameServerAllocation *v1.GameServerAllocation, opts metav1.CreateOptions) (result *v1.GameServerAllocation, err error) {
	result = &v1.GameServerAllocation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gameserverallocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gameServerAllocation).
		Do(ctx).
		Into(result)
	return
}
