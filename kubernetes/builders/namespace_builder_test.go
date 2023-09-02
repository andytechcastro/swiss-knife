package builders_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initNamespace() *builders.Namespace {
	namespace := builders.NewNamespaceBuilder("default")
	namespace.SetAnnotations(map[string]string{"my-first": "annotation"}).
		SetLabels(map[string]string{"my-first": "label"})
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
			"metadata": map[string]interface{}{
				"creationTimestamp": interface{}(nil),
				"name":              "default",
				"annotations": map[string]string{
					"my-first": "annotation",
				},
				"labels": map[string]string{
					"my-first": "label",
				},
			},
			"spec":   map[string]interface{}{},
			"status": map[string]interface{}{},
		},
	)
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlNS), string(yamlResult))
}
