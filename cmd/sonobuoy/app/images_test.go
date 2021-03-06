/*
Copyright the Sonobuoy contributors 2019

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

package app

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/vmware-tanzu/sonobuoy/pkg/image"
)

const e2eConfigContent = `---
buildImageRegistry: test-fake-registry.corp/fake-user
dockerGluster: test-fake-registry.corp/fake-user
dockerLibraryRegistry: test-fake-registry.corp/fake-user
e2eRegistry: test-fake-registry.corp/fake-user
e2eVolumeRegistry: test-fake-registry.corp/fake-user
gcRegistry: test-fake-registry.corp/fake-user
promoterE2eRegistry: test-fake-registry.corp/fake-user
sigStorageRegistry: test-fake-registry.corp/fake-user
`

func sampleE2eRegistryConfig() (string, error) {
	configFile, err := ioutil.TempFile("", "e2eRegistryConfig.yaml")
	if err != nil {
		return "", err
	}
	if _, err := configFile.Write([]byte(e2eConfigContent)); err != nil {
		return "", err
	}
	if err := configFile.Close(); err != nil {
		return "", err
	}
	return configFile.Name(), nil
}

func TestConvertImagesToPairs(t *testing.T) {
	configFileName, err := sampleE2eRegistryConfig()
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(configFileName)

	images := []string{
		"k8s.gcr.io/etcd:3.4.13-0",
	}
	expectedTagPairs := []image.TagPair{
		{
			Src: "k8s.gcr.io/etcd:3.4.13-0",
			Dst: "test-fake-registry.corp/fake-user/etcd:3.4.13-0",
		},
	}
	receivedTagPairs, err := convertImagesToPairs(images, "", configFileName, "v1.19.1")
	if err != nil {
		t.Error(err)
	}
	if receivedTagPairs[0] != expectedTagPairs[0] {
		t.Error(err)
	}
}
