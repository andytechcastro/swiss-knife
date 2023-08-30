package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

func initDeployment() *actions.Deployment {
	info := map[string]string{
		"service1": "nginx",
		"service2": "apache",
		"service3": "go",
		"service4": "docker",
	}
	objects := []runtime.Object{}
	for name, image := range info {
		deployment := builders.NewDeploymentBuilder()
		container := builders.NewContainerBuilder()
		container.SetName("testContainer").
			SetImage(image).
			SetTag("1").
			SetPort(80)
		buildedDeployment := deployment.SetName(name).
			SetNamespace("default").
			SetLabels(map[string]string{"test": "testing"}).
			SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
			SetReplicas(3).
			SetMatchLabels(map[string]string{"test": "testingmatch"}).
			AddContainer(*container.Build()).
			Build()
		objects = append(objects, buildedDeployment)
	}
	client := fake.NewSimpleClientset(objects...)
	objectsDynamic := []runtime.Object{}
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).Deployment
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
	deployment := builders.NewDeploymentBuilder()
	container := builders.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("java").
		SetTag("3").
		SetPort(80)
	buildedDeployment := deployment.SetName("service5").
		SetNamespace("default").
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		AddContainer(*container.Build()).
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
