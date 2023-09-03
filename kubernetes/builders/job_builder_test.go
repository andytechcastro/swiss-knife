package builders_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initJob() *builders.Job {
	container := builders.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("nginx").
		SetTag("1").
		SetPort(80)
	pod := builders.NewPodBuilder("test")
	pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
	job := builders.NewJobBuilder("test")
	job.SetPodTemplate(*pod.BuildTemplate()).
		SetNamespace("testNamespace").
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		SetMatchLabels(map[string]string{"test": "testingmatch"}).
		SetBackOffLimit(10).
		SetTTLSecondsAfterFinished(200)
	return job
}

func TestBuildJob(t *testing.T) {
	job := initJob()
	buildedJob := job.Build()
	assert.Equal(t, "test", buildedJob.Name)
	assert.Equal(t, "testNamespace", buildedJob.Namespace)
	assert.Equal(t, map[string]string{"test": "testing"}, buildedJob.Labels)
	assert.Equal(t, map[string]string{"annotation": "testAnnotation"}, buildedJob.Annotations)
	assert.Equal(t, int32(10), *buildedJob.Spec.BackoffLimit)
	assert.Equal(t, int32(200), *buildedJob.Spec.TTLSecondsAfterFinished)
	assert.Equal(t, "testContainer", buildedJob.Spec.Template.Spec.Containers[0].Name)
	assert.Equal(t, "nginx:1", buildedJob.Spec.Template.Spec.Containers[0].Image)
	assert.Equal(t, int32(80), buildedJob.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort)
}

func TestJobToYaml(t *testing.T) {
	job := initJob()
	job.Build()
	yamlJob := job.ToYaml()
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
				"backoffLimit":            10,
				"ttlSecondsAfterFinished": 200,
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
			"status": map[string]interface{}{},
		},
	)
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlResult), string(yamlJob))
}
