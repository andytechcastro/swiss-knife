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

func initCronJob() *actions.CronJob {
	info := map[string]string{
		"cronJob1": "nginx",
		"cronJob2": "apache",
		"cronJob3": "go",
		"cronJob4": "docker",
	}
	objects := []runtime.Object{}
	for name, image := range info {
		cronJob := batchv1.NewCronJobBuilder(name)
		container := builders.NewContainerBuilder()
		container.SetName("testContainer").
			SetImage(image).
			SetTag("1").
			SetPort(80)
		pod := corev1.NewPodBuilder(name)
		pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
		job := batchv1.NewJobBuilder("test")
		job.SetPodTemplate(*pod.BuildTemplate()).
			SetMatchLabels(map[string]string{"test": "testingmatch"}).
			SetBackOffLimit(10).
			SetTTLSecondsAfterFinished(200)
		buildedCronJob := cronJob.SetNamespace("default").
			SetLabels(map[string]string{"test": "testing"}).
			SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
			SetJobTemplate(*job.BuildTemplate()).
			Build()
		objects = append(objects, buildedCronJob)
	}
	client := fake.NewSimpleClientset(objects...)
	objectsDynamic := []runtime.Object{}
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).CronJob
	actions.Namespace("default")
	return actions
}

func TestGetCronJob(t *testing.T) {
	actions := initCronJob()
	cronJob, _ := actions.Get("cronJob3")
	assert.Equal(t, "cronJob3", cronJob.Name)
	assert.Equal(t, "go:1", cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image)
}

func TestCreateCronJob(t *testing.T) {
	actions := initCronJob()
	cronJob := batchv1.NewCronJobBuilder("cronJob5")
	container := builders.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("java").
		SetTag("3").
		SetPort(80)
	pod := corev1.NewPodBuilder("cronJob5")
	pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
	job := batchv1.NewJobBuilder("test")
	job.SetPodTemplate(*pod.BuildTemplate()).
		SetMatchLabels(map[string]string{"test": "testingmatch"}).
		SetBackOffLimit(10).
		SetTTLSecondsAfterFinished(200)
	buildedCronJob := cronJob.SetNamespace("default").
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		SetJobTemplate(*job.BuildTemplate()).
		Build()
	actions.Create(buildedCronJob)
	newCronJob, _ := actions.Get("cronJob5")
	cronJobs, _ := actions.List()
	assert.Equal(t, "cronJob5", newCronJob.Name)
	assert.Equal(t, "java:3", newCronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image)
	assert.Equal(t, 5, len(cronJobs.Items))
}

func TestUpdateCronJob(t *testing.T) {
	actions := initCronJob()
	cronJob, _ := actions.Get("cronJob3")
	cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image = "go:1.21"
	actions.Update(cronJob)
	updatedCronJob, _ := actions.Get("cronJob3")
	assert.Equal(t, "go:1.21", updatedCronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image)
}

func TestDeleteCronJob(t *testing.T) {
	actions := initCronJob()
	actions.Delete("cronJob4")
	cronJobs, _ := actions.List()
	assert.Equal(t, 3, len(cronJobs.Items))
	for _, cronJob := range cronJobs.Items {
		assert.NotEqual(t, "cronJob4", cronJob.Name)
	}
}

func TestListCronJob(t *testing.T) {
	actions := initCronJob()
	cronJobs, _ := actions.List()
	assert.Equal(t, 4, len(cronJobs.Items))
}
