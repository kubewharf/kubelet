/*
Copyright 2026 The Kubernetes Authors.

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

package v1alpha1

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestResourcePluginOptionsExtraResources(t *testing.T) {
	want := []string{"resource.example.com/foo", "resource.example.com/bar"}
	options := &ResourcePluginOptions{ExtraResources: want}

	data, err := proto.Marshal(options)
	if err != nil {
		t.Fatalf("marshal ResourcePluginOptions: %v", err)
	}

	got := &ResourcePluginOptions{}
	if err := proto.Unmarshal(data, got); err != nil {
		t.Fatalf("unmarshal ResourcePluginOptions: %v", err)
	}

	if !proto.Equal(options, got) {
		t.Fatalf("round-trip mismatch: got %v, want %v", got.GetExtraResources(), want)
	}
}
