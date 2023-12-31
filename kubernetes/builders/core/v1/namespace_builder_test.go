package builders_test

import (
	"testing"

	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initNamespace() *corev1.Namespace {
	namespace := corev1.NewNamespaceBuilder("default")
	namespace.SetAnnotations(map[string]string{"my-first": "annotation"}).
		SetLabels(map[string]string{"my-first": "label"})
	return namespace
}

func TestBuildNamespace(t *testing.T) {
	namespace := initNamespace()
	buildedNamespace := namespace.Build()
	assert.Equal(t, "default", buildedNamespace.Name)
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
	assert.YAMLEq(t, string(yamlResult), string(yamlNS))
}
