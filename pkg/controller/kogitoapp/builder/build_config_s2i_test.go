// Copyright 2019 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builder

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"testing"

	buildv1 "github.com/openshift/api/build/v1"
)

func Test_getBCS2ILimitsAsIntString(t *testing.T) {
	type args struct {
		buildConfig *buildv1.BuildConfig
	}
	var tests = []struct {
		name            string
		args            args
		wantLimitCPU    string
		wantLimitMemory string
	}{
		{"With Limits", args{buildConfig: &buildv1.BuildConfig{
			Spec: buildv1.BuildConfigSpec{
				CommonSpec: buildv1.CommonSpec{
					Resources: v1.ResourceRequirements{
						Limits: map[v1.ResourceName]resource.Quantity{
							v1.ResourceCPU:    resource.MustParse("1"),
							v1.ResourceMemory: resource.MustParse("512Mi"),
						},
					},
				},
			},
		}}, "1", "536870912"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLimitCPU, gotLimitMemory := getBCS2ILimitsAsIntString(tt.args.buildConfig)
			if gotLimitCPU != tt.wantLimitCPU {
				t.Errorf("getBCS2ILimitsAsIntString() gotLimitCPU = %v, want %v", gotLimitCPU, tt.wantLimitCPU)
			}
			if gotLimitMemory != tt.wantLimitMemory {
				t.Errorf("getBCS2ILimitsAsIntString() gotLimitMemory = %v, want %v", gotLimitMemory, tt.wantLimitMemory)
			}
		})
	}
}