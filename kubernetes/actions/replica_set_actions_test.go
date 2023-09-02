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

func initReplicaSet() *actions.ReplicaSet {
	info := map[string]string{
		"replicaset1": "nginx",
		"replicaset2": "apache",
		"replicaset3": "go",
		"replicaset4": "docker",
	}
	objects := []runtime.Object{}
	for name, image := range info {
		replicaSet := builders.NewReplicaSetBuilder(name)
		container := builders.NewContainerBuilder()
		container.SetName("testContainer").
			SetImage(image).
			SetTag("1").
			SetPort(80)
		pod := builders.NewPodBuilder(name)
		pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
		buildedReplicaSet := replicaSet.SetNamespace("default").
			SetLabels(map[string]string{"test": "testing"}).
			SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
			SetReplicas(3).
			SetMatchLabels(map[string]string{"test": "testingmatch"}).
			SetPodTemplate(*pod.BuildTemplate()).
			Build()
		objects = append(objects, buildedReplicaSet)
	}
	client := fake.NewSimpleClientset(objects...)
	objectsDynamic := []runtime.Object{}
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).ReplicaSet
	actions.Namespace("default")
	return actions
}

func TestGetReplicaSet(t *testing.T) {
	actions := initReplicaSet()
	replicaSet, _ := actions.Get("replicaset3")
	assert.Equal(t, "replicaset3", replicaSet.Name)
	assert.Equal(t, "go:1", replicaSet.Spec.Template.Spec.Containers[0].Image)
}

func TestCreateReplicaSet(t *testing.T) {
	actions := initReplicaSet()
	replicaSet := builders.NewReplicaSetBuilder("replicaset5")
	container := builders.NewContainerBuilder()
	container.SetName("testContainer").
		SetImage("java").
		SetTag("3").
		SetPort(80)
	pod := builders.NewPodBuilder("replicaset5")
	pod.SetLabels(map[string]string{"test": "testingmatch"}).AddContainer(*container.Build())
	buildedReplicaSet := replicaSet.SetNamespace("default").
		SetLabels(map[string]string{"test": "testing"}).
		SetAnnotations(map[string]string{"annotation": "testAnnotation"}).
		SetPodTemplate(*pod.BuildTemplate()).
		Build()
	actions.Create(buildedReplicaSet)
	newReplicaSet, _ := actions.Get("replicaset5")
	replicaSets, _ := actions.List()
	assert.Equal(t, "replicaset5", newReplicaSet.Name)
	assert.Equal(t, "java:3", newReplicaSet.Spec.Template.Spec.Containers[0].Image)
	assert.Equal(t, 5, len(replicaSets.Items))
}

func TestUpdateReplicaSet(t *testing.T) {
	actions := initReplicaSet()
	replicaSet, _ := actions.Get("replicaset3")
	replicaSet.Spec.Template.Spec.Containers[0].Image = "go:1.21"
	actions.Update(replicaSet)
	updatedReplicaSet, _ := actions.Get("replicaset3")
	assert.Equal(t, "go:1.21", updatedReplicaSet.Spec.Template.Spec.Containers[0].Image)
}

func TestDeleteReplicaSet(t *testing.T) {
	actions := initReplicaSet()
	actions.Delete("replicaset4")
	replicaSets, _ := actions.List()
	assert.Equal(t, 3, len(replicaSets.Items))
	for _, replicaSet := range replicaSets.Items {
		assert.NotEqual(t, "replicaset4", replicaSet.Name)
	}
}

func TestListReplicaSet(t *testing.T) {
	actions := initReplicaSet()
	replicaSets, _ := actions.List()
	assert.Equal(t, 4, len(replicaSets.Items))
}
