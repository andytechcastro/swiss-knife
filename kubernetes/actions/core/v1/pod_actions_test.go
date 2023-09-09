package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	actionsCoreV1 "github.com/andytechcastro/swiss-knife/kubernetes/actions/core/v1"
	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

func initPod() *actionsCoreV1.Pod {
	info := map[string]string{
		"service1": "nginx",
		"service2": "apache",
		"service3": "go",
		"service4": "docker",
	}
	objects := []runtime.Object{}
	for name, image := range info {
		pod := corev1.NewPodBuilder(name)
		container := corev1.NewContainerBuilder()
		container.SetName("testContainer").
			SetImage(image).
			SetTag("1").
			SetPort(80)
		buildedPod := pod.SetNamespace("default").
			SetLabels(map[string]string{"test": "testing"}).
			SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
			AddContainer(*container.Build()).
			Build()
		objects = append(objects, buildedPod)
	}
	objectsDynamic := []runtime.Object{}
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	client := fake.NewSimpleClientset(objects...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).CoreV1.Pod
	actions.Namespace("default")
	return actions
}

func TestGetPod(t *testing.T) {
	actions := initPod()
	pod, _ := actions.Get("service3")
	assert.Equal(t, "service3", pod.Name)
	assert.Equal(t, "go:1", pod.Spec.Containers[0].Image)
}

func TestCreatePod(t *testing.T) {
	actions := initPod()
	pod := corev1.NewPodBuilder("service5")
	container := corev1.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("java").
		SetTag("3").
		SetPort(80)
	buildedPod := pod.SetNamespace("default").
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		AddContainer(*container.Build()).
		Build()
	actions.Create(buildedPod)
	newPod, _ := actions.Get("service5")
	pods, _ := actions.List()
	assert.Equal(t, "service5", newPod.Name)
	assert.Equal(t, "java:3", newPod.Spec.Containers[0].Image)
	assert.Equal(t, 5, len(pods.Items))
}

func TestUpdatePod(t *testing.T) {
	actions := initPod()
	pod, _ := actions.Get("service3")
	pod.Spec.Containers[0].Image = "go:1.21"
	actions.Update(pod)
	updatedPod, _ := actions.Get("service3")
	assert.Equal(t, "go:1.21", updatedPod.Spec.Containers[0].Image)
}

func TestDeletePod(t *testing.T) {
	actions := initPod()
	actions.Delete("service4")
	pods, _ := actions.List()
	assert.Equal(t, 3, len(pods.Items))
	for _, pod := range pods.Items {
		assert.NotEqual(t, "service4", pod.Name)
	}
}

func TestListPod(t *testing.T) {
	actions := initPod()
	pods, _ := actions.List()
	assert.Equal(t, 4, len(pods.Items))
}
