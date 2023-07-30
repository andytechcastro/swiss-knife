package builders_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	kube "github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initNamespace() *builders.Namespace {
	namespace := kube.NewNamespaceBuilder()
	namespace.SetName("default")
	return namespace
}

func TestBuildNamespace(t *testing.T) {
	namespace := initNamespace()
	buildedNamespace := namespace.Build()
	assert.Equal(t, buildedNamespace.Name, "default")
}

func TestNamespaceToYaml(t *testing.T) {
	namespace := initNamespace()
	namespace.Build()
	yamlNS := namespace.ToYaml()
	interfaceResult := map[string]interface{}(
		map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Namespace",
			"metadata": map[string]interface{}{
				"creationTimestamp": interface{}(nil),
				"name":              "default",
			},
			"spec":   map[string]interface{}{},
			"status": map[string]interface{}{},
		},
	)
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlNS), string(yamlResult))
}
