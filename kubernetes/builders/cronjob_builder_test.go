package builders_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initCronJob() *builders.CronJob {
	container := builders.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("nginx").
		SetTag("1").
		SetPort(80)
	pod := builders.NewPodBuilder("test")
	pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
	job := builders.NewJobBuilder("test")
	job.SetPodTemplate(*pod.BuildTemplate()).
		SetMatchLabels(map[string]string{"test": "testingmatch"}).
		SetBackOffLimit(10).
		SetTTLSecondsAfterFinished(200)
	cronJob := builders.NewCronJobBuilder("test")
	cronJob.SetJobTemplate(*job.BuildTemplate()).
		SetNamespace("testNamespace").
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		SetSchedule("* * * * *")
	return cronJob
}

func TestBuildCronJob(t *testing.T) {
	cronJob := initCronJob()
	buildedCronJob := cronJob.Build()
	assert.Equal(t, "test", buildedCronJob.Name)
	assert.Equal(t, "testNamespace", buildedCronJob.Namespace)
	assert.Equal(t, map[string]string{"test": "testing"}, buildedCronJob.Labels)
	assert.Equal(t, map[string]string{"annotation": "testAnnotation"}, buildedCronJob.Annotations)
	assert.Equal(t, "* * * * *", *&buildedCronJob.Spec.Schedule)
	assert.Equal(t, int32(10), *buildedCronJob.Spec.JobTemplate.Spec.BackoffLimit)
	assert.Equal(t, int32(200), *buildedCronJob.Spec.JobTemplate.Spec.TTLSecondsAfterFinished)
	assert.Equal(t, "testContainer", buildedCronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Name)
	assert.Equal(t, "nginx:1", buildedCronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image)
	assert.Equal(t, int32(80), buildedCronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort)
}

func TestCronJobToYaml(t *testing.T) {
	cronJob := initCronJob()
	cronJob.Build()
	yamlCronJob := cronJob.ToYaml()
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
				"schedule": "* * * * *",
				"jobTemplate": map[string]interface{}{
					"metadata": map[string]interface{}{
						"creationTimestamp": interface{}(nil),
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
				},
			},
			"status": map[string]interface{}{},
		},
	)
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlResult), string(yamlCronJob))
}
