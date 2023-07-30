package builders_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	kube "github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initDeployment() *builders.Deployment {
	container := kube.NewContainerBuilder()
	container.SetName("testContainer")
	container.SetImage("nginx")
	container.SetTag("1")
	container.SetPort(80)
	deployment := kube.NewDeploymentBuilder()
	deployment.AddContainer(container.Build())
	deployment.SetName("test")
	deployment.SetNamespace("testNamespace")
	deployment.SetReplicas(3)
	deployment.SetLabels(map[string]string{"test": "testing"})
	deployment.SetAnnotations(map[string]string{"annotation": "testAnnotation"})
	deployment.SetMatchLabels(map[string]string{"test": "testingmatch"})
	return deployment
}

func TestBuildDeployment(t *testing.T) {
	deployment := initDeployment()
	buildedDeployment := deployment.Build()
	assert.Equal(t, buildedDeployment.Name, "test")
	assert.Equal(t, buildedDeployment.Namespace, "testNamespace")
	assert.Equal(t, buildedDeployment.Labels, map[string]string{"test": "testing"})
	assert.Equal(t, buildedDeployment.Annotations, map[string]string{"annotation": "testAnnotation"})
	assert.Equal(t, *buildedDeployment.Spec.Replicas, int32(3))
	assert.Equal(t, buildedDeployment.Spec.Template.Spec.Containers[0].Name, "testContainer")
	assert.Equal(t, buildedDeployment.Spec.Template.Spec.Containers[0].Image, "nginx:1")
	assert.Equal(t, buildedDeployment.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort, int32(80))
}

func TestDeploymentToYaml(t *testing.T) {
	deployment := initDeployment()
	deployment.Build()
	yamlDeploy := deployment.ToYaml()
	interfaceResult := map[string]interface{}(
		map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
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
	assert.YAMLEq(t, string(yamlDeploy), string(yamlResult))
}
