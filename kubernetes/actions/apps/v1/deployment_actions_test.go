package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	actionsAppsV1 "github.com/andytechcastro/swiss-knife/kubernetes/actions/apps/v1"
	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	appsv1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/apps/v1"
	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

func initDeployment() *actionsAppsV1.Deployment {
	info := map[string]string{
		"service1": "nginx",
		"service2": "apache",
		"service3": "go",
		"service4": "docker",
	}
	objects := []runtime.Object{}
	for name, image := range info {
		deployment := appsv1.NewDeploymentBuilder(name)
		container := builders.NewContainerBuilder()
		container.SetName("testContainer").
			SetImage(image).
			SetTag("1").
			SetPort(80)
		pod := corev1.NewPodBuilder(name)
		pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
		buildedDeployment := deployment.SetNamespace("default").
			SetLabels(map[string]string{"test": "testing"}).
			SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
			SetReplicas(3).
			SetMatchLabels(map[string]string{"test": "testingmatch"}).
			SetPodTemplate(*pod.BuildTemplate()).
			Build()
		objects = append(objects, buildedDeployment)
	}
	client := fake.NewSimpleClientset(objects...)
	objectsDynamic := []runtime.Object{}
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).AppsV1.Deployment
	actions.Namespace("default")
	return actions
}

func TestGetDeployment(t *testing.T) {
	actions := initDeployment()
	deployment, _ := actions.Get("service3")
	assert.Equal(t, "service3", deployment.Name)
	assert.Equal(t, "go:1", deployment.Spec.Template.Spec.Containers[0].Image)
}

func TestCreateDeployment(t *testing.T) {
	actions := initDeployment()
	deployment := appsv1.NewDeploymentBuilder("service5")
	container := builders.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("java").
		SetTag("3").
		SetPort(80)
	pod := corev1.NewPodBuilder("service5")
	pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
	buildedDeployment := deployment.SetNamespace("default").
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		SetPodTemplate(*pod.BuildTemplate()).
		Build()
	actions.Create(buildedDeployment)
	newDeployment, _ := actions.Get("service5")
	deployments, _ := actions.List()
	assert.Equal(t, "service5", newDeployment.Name)
	assert.Equal(t, "java:3", newDeployment.Spec.Template.Spec.Containers[0].Image)
	assert.Equal(t, 5, len(deployments.Items))
}

func TestUpdateDeployment(t *testing.T) {
	actions := initDeployment()
	deployment, _ := actions.Get("service3")
	deployment.Spec.Template.Spec.Containers[0].Image = "go:1.21"
	actions.Update(deployment)
	updatedDeployment, _ := actions.Get("service3")
	assert.Equal(t, "go:1.21", updatedDeployment.Spec.Template.Spec.Containers[0].Image)
}

func TestDeleteDeployment(t *testing.T) {
	actions := initDeployment()
	actions.Delete("service4")
	deployments, _ := actions.List()
	assert.Equal(t, 3, len(deployments.Items))
	for _, deployment := range deployments.Items {
		assert.NotEqual(t, "service4", deployment.Name)
	}
}

func TestListDeployment(t *testing.T) {
	actions := initDeployment()
	deployments, _ := actions.List()
	assert.Equal(t, 4, len(deployments.Items))
}
