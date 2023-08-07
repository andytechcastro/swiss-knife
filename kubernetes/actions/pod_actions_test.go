package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func initPod() *actions.Actions {
	info := map[string]string{
		"service1": "nginx",
		"service2": "apache",
		"service3": "go",
		"service4": "docker",
	}
	objects := []runtime.Object{}
	for name, image := range info {
		pod := builders.NewPodBuilder()
		container := builders.NewContainerBuilder()
		container.SetName("testContainer").
			SetImage(image).
			SetTag("1").
			SetPort(80)
		buildedPod := pod.SetName(name).
			SetNamespace("default").
			SetLabels(map[string]string{"test": "testing"}).
			SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
			AddContainer(*container.Build()).
			Build()
		objects = append(objects, buildedPod)
	}
	client := fake.NewSimpleClientset(objects...)
	action := actions.NewActions(client)
	return action
}

func TestGetPod(t *testing.T) {
	actions := initPod()
	pod, _ := actions.GetPod("service3")
	assert.Equal(t, "service3", pod.Name)
	assert.Equal(t, "go:1", pod.Spec.Containers[0].Image)
}

func TestCreatePod(t *testing.T) {
	actions := initPod()
	pod := builders.NewPodBuilder()
	container := builders.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("java").
		SetTag("3").
		SetPort(80)
	buildedPod := pod.SetName("service5").
		SetNamespace("default").
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		AddContainer(*container.Build()).
		Build()
	actions.CreatePod(buildedPod)
	newPod, _ := actions.GetPod("service5")
	pods, _ := actions.ListPod()
	assert.Equal(t, "service5", newPod.Name)
	assert.Equal(t, "java:3", newPod.Spec.Containers[0].Image)
	assert.Equal(t, 5, len(pods.Items))
}

func TestUpdatePod(t *testing.T) {
	actions := initPod()
	pod, _ := actions.GetPod("service3")
	pod.Spec.Containers[0].Image = "go:1.21"
	actions.UpdatePod(pod)
	updatedPod, _ := actions.GetPod("service3")
	assert.Equal(t, "go:1.21", updatedPod.Spec.Containers[0].Image)
}

func TestDeletePod(t *testing.T) {
	actions := initPod()
	actions.DeletePod("service4")
	pods, _ := actions.ListPod()
	assert.Equal(t, 3, len(pods.Items))
	for _, pod := range pods.Items {
		assert.NotEqual(t, "service4", pod.Name)
	}
}

func TestListPod(t *testing.T) {
	actions := initPod()
	pods, _ := actions.ListPod()
	assert.Equal(t, 4, len(pods.Items))
}
