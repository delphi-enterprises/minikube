/*
Copyright 2019 The Kubernetes Authors All rights reserved.

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

package lock

import "testing"

func TestUserMutexSpec(t *testing.T) {
	forceID = "test"

	var tests = []struct {
		description string
		path        string
	}{
		{
			description: "standard",
			path:        "/foo/bar",
		},
		{
			description: "deep directory",
			path:        "/foo/bar/baz/bat",
		},
		{
			description: "underscores",
			path:        "/foo_bar/baz",
		},
		{
			description: "starts with number",
			path:        "/foo/2bar/baz",
		},
		{
			description: "starts with punctuation",
			path:        "/.foo/bar",
		},
		{
			description: "long filename",
			path:        "/very-very-very-very-very-very-very-very-long/bar",
		},
		{
			description: "Windows kubeconfig",
			path:        `C:\Users\admin/.kube/config`,
		},
		{
			description: "Windows json",
			path:        `C:\Users\admin\.minikube\profiles\containerd-20191210T212325.7356633-8584\config.json`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := UserMutexSpec(tc.path)
			if len(got.Name) > 40 {
				t.Errorf("%s mutex name is too long", got.Name)
			}
		})
	}
}
