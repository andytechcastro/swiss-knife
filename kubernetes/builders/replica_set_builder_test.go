package builders_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	appsv1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/apps/v1"
	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initReplicaSet() *appsv1.ReplicaSet {
	container := builders.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("nginx").
		SetTag("1").
		SetPort(80)
	pod := corev1.NewPodBuilder("test")
	pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
	replicaSet := appsv1.NewReplicaSetBuilder("test")
	replicaSet.SetPodTemplate(*pod.BuildTemplate()).
		SetNamespace("testNamespace").
		SetReplicas(3).
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		SetMatchLabels(map[string]string{"test": "testingmatch"})
	return replicaSet
}

func TestBuildReplicaSet(t *testing.T) {
	replicaSet := initReplicaSet()
	buildedReplicaSet := replicaSet.Build()
	assert.Equal(t, "test", buildedReplicaSet.Name)
	assert.Equal(t, "testNamespace", buildedReplicaSet.Namespace)
	assert.Equal(t, map[string]string{"test": "testing"}, buildedReplicaSet.Labels)
	assert.Equal(t, map[string]string{"annotation": "testAnnotation"}, buildedReplicaSet.Annotations)
	assert.Equal(t, int32(3), *buildedReplicaSet.Spec.Replicas)
	assert.Equal(t, "testContainer", buildedReplicaSet.Spec.Template.Spec.Containers[0].Name)
	assert.Equal(t, "nginx:1", buildedReplicaSet.Spec.Template.Spec.Containers[0].Image)
	assert.Equal(t, int32(80), buildedReplicaSet.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort)
}

func TestReplicaSetToYaml(t *testing.T) {
	replicaSet := initReplicaSet()
	replicaSet.Build()
	yamlReplicaSet := replicaSet.ToYaml()
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
				"replicas": 3,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"test": "testingmatch",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"creationTimestamp": interface{}(nil),
						"labels": map[string]interface{}{
							"test": "testingmatch",
						},
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
				},
			},
			"status": map[string]interface{}{
				"replicas": 0,
			},
		},
	)
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlResult), string(yamlReplicaSet))
}
