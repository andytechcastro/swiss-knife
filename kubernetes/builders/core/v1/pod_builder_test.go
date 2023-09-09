package builders_test

import (
	"testing"

	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initPod() *corev1.Pod {
	container := corev1.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("nginx").
		SetTag("1").
		SetPort(80)
	pod := corev1.NewPodBuilder("test")
	pod.AddContainer(*container.Build()).
		SetNamespace("testNamespace").
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"})
	return pod
}

func TestBuildPod(t *testing.T) {
	pod := initPod()
	buildedPod := pod.Build()
	assert.Equal(t, buildedPod.Name, "test")
	assert.Equal(t, buildedPod.Namespace, "testNamespace")
	assert.Equal(t, buildedPod.Labels, map[string]string{"test": "testing"})
	assert.Equal(t, buildedPod.Annotations, map[string]string{"annotation": "testAnnotation"})
	assert.Equal(t, buildedPod.Spec.Containers[0].Name, "testContainer")
	assert.Equal(t, buildedPod.Spec.Containers[0].Image, "nginx:1")
	assert.Equal(t, buildedPod.Spec.Containers[0].Ports[0].ContainerPort, int32(80))
}

func TestBuildPodTemplate(t *testing.T) {
	pod := initPod()
	buildedPod := pod.BuildTemplate()
	assert.Equal(t, buildedPod.Name, "")
	assert.Equal(t, buildedPod.Namespace, "")
	assert.Equal(t, buildedPod.Labels, map[string]string{"test": "testing"})
	assert.Equal(t, buildedPod.Annotations, map[string]string{"annotation": "testAnnotation"})
	assert.Equal(t, buildedPod.Spec.Containers[0].Name, "testContainer")
	assert.Equal(t, buildedPod.Spec.Containers[0].Image, "nginx:1")
	assert.Equal(t, buildedPod.Spec.Containers[0].Ports[0].ContainerPort, int32(80))
}

func TestPodToYaml(t *testing.T) {
	pod := initPod()
	pod.Build()
	yamlPod := pod.ToYaml()
	interfaceResult := map[string]interface{}(
		map[string]interface{}{
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"annotation": "testAnnotation",
				},
				"creationTimestamp": interface{}(nil),
				"labels": map[string]interface{}{
					"test": "testing",
				},
				"name":      "test",
				"namespace": "testNamespace",
			},
			"spec": map[string]interface{}{
				"containers": []interface{}{
					map[string]interface{}{
						"image": "nginx:1",
						"name":  "testContainer",
						"ports": []interface{}{
							map[string]interface{}{
								"containerPort": 80,
								"name":          "http",
								"protocol":      "TCP",
							},
						},
						"resources": map[string]interface{}{},
					},
				},
			},
			"status": map[string]interface{}{},
		},
	)
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlPod), string(yamlResult))
}
