package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	batchv1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/batch/v1"
	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

func initJob() *actions.Job {
	info := map[string]string{
		"job1": "nginx",
		"job2": "apache",
		"job3": "go",
		"job4": "docker",
	}
	objects := []runtime.Object{}
	for name, image := range info {
		job := batchv1.NewJobBuilder(name)
		container := builders.NewContainerBuilder()
		container.SetName("testContainer").
			SetImage(image).
			SetTag("1").
			SetPort(80)
		pod := corev1.NewPodBuilder(name)
		pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
		buildedJob := job.SetNamespace("default").
			SetLabels(map[string]string{"test": "testing"}).
			SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
			SetMatchLabels(map[string]string{"test": "testingmatch"}).
			SetPodTemplate(*pod.BuildTemplate()).
			Build()
		objects = append(objects, buildedJob)
	}
	client := fake.NewSimpleClientset(objects...)
	objectsDynamic := []runtime.Object{}
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).Job
	actions.Namespace("default")
	return actions
}

func TestGetJob(t *testing.T) {
	actions := initJob()
	job, _ := actions.Get("job3")
	assert.Equal(t, "job3", job.Name)
	assert.Equal(t, "go:1", job.Spec.Template.Spec.Containers[0].Image)
}

func TestCreateJob(t *testing.T) {
	actions := initJob()
	job := batchv1.NewJobBuilder("job5")
	container := builders.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("java").
		SetTag("3").
		SetPort(80)
	pod := corev1.NewPodBuilder("job5")
	pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
	buildedJob := job.SetNamespace("default").
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		SetPodTemplate(*pod.BuildTemplate()).
		Build()
	actions.Create(buildedJob)
	newJob, _ := actions.Get("job5")
	jobs, _ := actions.List()
	assert.Equal(t, "job5", newJob.Name)
	assert.Equal(t, "java:3", newJob.Spec.Template.Spec.Containers[0].Image)
	assert.Equal(t, 5, len(jobs.Items))
}

func TestUpdateJob(t *testing.T) {
	actions := initJob()
	job, _ := actions.Get("job3")
	job.Spec.Template.Spec.Containers[0].Image = "go:1.21"
	actions.Update(job)
	updatedJob, _ := actions.Get("job3")
	assert.Equal(t, "go:1.21", updatedJob.Spec.Template.Spec.Containers[0].Image)
}

func TestDeleteJob(t *testing.T) {
	actions := initJob()
	actions.Delete("job4")
	jobs, _ := actions.List()
	assert.Equal(t, 3, len(jobs.Items))
	for _, job := range jobs.Items {
		assert.NotEqual(t, "job4", job.Name)
	}
}

func TestListJob(t *testing.T) {
	actions := initJob()
	jobs, _ := actions.List()
	assert.Equal(t, 4, len(jobs.Items))
}
