package builders_test

import (
	"testing"

	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initConfigMap() *corev1.ConfigMap {
	configmap := corev1.NewConfigMapBuilder("configmap")
	configmap.SetNamespace("default").
		SetData(map[string]string{
			"myKey": "myValue",
		})
	return configmap
}

func TestBuildConfigMap(t *testing.T) {
	configmap := initConfigMap()
	buildedConfigMap := configmap.Build()
	assert.Equal(t, "configmap", buildedConfigMap.Name)
}

func TestConfigMapToYaml(t *testing.T) {
	configmap := initConfigMap()
	configmap.SetImmutable()
	configmap.AddData("secondKey", "secondValue")
	configmap.Build()
	yamlCM := configmap.ToYaml()
	interfaceResult := map[string]interface{}(
		map[string]interface{}{
			"metadata": map[string]interface{}{
				"creationTimestamp": interface{}(nil),
				"name":              "configmap",
				"namespace":         "default",
			},
			"data": map[string]interface{}{
				"myKey":     "myValue",
				"secondKey": "secondValue",
			},
			"immutable": true,
		},
	)
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlCM), string(yamlResult))
}
