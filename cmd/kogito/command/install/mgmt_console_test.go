// Copyright 2020 Red Hat, Inc. and/or its affiliates
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

package install

import (
	"fmt"
	"testing"

	"github.com/kiegroup/kogito-cloud-operator/cmd/kogito/command/context"
	"github.com/kiegroup/kogito-cloud-operator/cmd/kogito/command/test"
	"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1"
	"github.com/kiegroup/kogito-cloud-operator/pkg/client/kubernetes"
	"github.com/kiegroup/kogito-cloud-operator/pkg/infrastructure"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_DeployMgmtConsoleCmd_DefaultConfiguration(t *testing.T) {
	ns := t.Name()
	cli := fmt.Sprintf("install mgmt-console -p %s", ns)
	ctx := test.SetupCliTest(cli, context.CommandFactory{BuildCommands: BuildCommands}, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}})
	lines, _, err := test.ExecuteCli()

	assert.NoError(t, err)
	assert.Contains(t, lines, "Kogito Management Console Service successfully installed")

	mgmtConsole := &v1alpha1.KogitoMgmtConsole{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      infrastructure.DefaultMgmtConsoleName,
		},
	}
	exits, err := kubernetes.ResourceC(ctx.Client).Fetch(mgmtConsole)
	assert.NoError(t, err)
	assert.True(t, exits)
	assert.Equal(t, mgmtConsole.Spec.Image.Name, "")
	assert.False(t, mgmtConsole.Spec.InsecureImageRegistry)
}

func Test_DeployMgmtConsoleCmd_CustomConfiguration(t *testing.T) {
	ns := t.Name()
	cli := fmt.Sprintf("install mgmt-console -p %s --image docker.io/project/mgmt-console:latest --insecure-image-registry --http-port 9090", ns)
	ctx := test.SetupCliTest(cli, context.CommandFactory{BuildCommands: BuildCommands}, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}})
	lines, _, err := test.ExecuteCli()

	assert.NoError(t, err)
	assert.Contains(t, lines, "Kogito Management Console Service successfully installed")

	mgmtConsole := &v1alpha1.KogitoMgmtConsole{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      infrastructure.DefaultMgmtConsoleName,
		},
	}
	exits, err := kubernetes.ResourceC(ctx.Client).Fetch(mgmtConsole)
	assert.NoError(t, err)
	assert.True(t, exits)
	assert.True(t, mgmtConsole.Spec.InsecureImageRegistry)
	assert.Equal(t, "mgmt-console", mgmtConsole.Spec.Image.Name)
	assert.Equal(t, "project", mgmtConsole.Spec.Image.Namespace)
	assert.Equal(t, "docker.io", mgmtConsole.Spec.Image.Domain)
	assert.Equal(t, "latest", mgmtConsole.Spec.Image.Tag)
	assert.Equal(t, int32(9090), mgmtConsole.Spec.HTTPPort)

}
