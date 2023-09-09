package builders_test

import (
	"testing"

	appsv1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/apps/v1"
	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initDeployment() *appsv1.Deployment {
	container := corev1.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("nginx").
		SetTag("1").
		SetPort(80)
	pod := corev1.NewPodBuilder("test")
	pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
	deployment := appsv1.NewDeploymentBuilder("test")
	deployment.SetPodTemplate(*pod.BuildTemplate()).
		SetNamespace("testNamespace").
		SetReplicas(3).
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		SetMatchLabels(map[string]string{"test": "testingmatch"})
	return deployment
}

func TestBuildDeployment(t *testing.T) {
	deployment := initDeployment()
	buildedDeployment := deployment.Build()
	assert.Equal(t, "test", buildedDeployment.Name)
	assert.Equal(t, "testNamespace", buildedDeployment.Namespace)
	assert.Equal(t, map[string]string{"test": "testing"}, buildedDeployment.Labels)
	assert.Equal(t, map[string]string{"annotation": "testAnnotation"}, buildedDeployment.Annotations)
	assert.Equal(t, int32(3), *buildedDeployment.Spec.Replicas)
	assert.Equal(t, "testContainer", buildedDeployment.Spec.Template.Spec.Containers[0].Name)
	assert.Equal(t, "nginx:1", buildedDeployment.Spec.Template.Spec.Containers[0].Image)
	assert.Equal(t, int32(80), buildedDeployment.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort)
}

func TestDeploymentToYaml(t *testing.T) {
	deployment := initDeployment()
	deployment.Build()
	yamlDeploy := deployment.ToYaml()
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
				"strategy": map[string]interface{}{},
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
			"status": map[string]interface{}{},
		},
	)
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlResult), string(yamlDeploy))
}
