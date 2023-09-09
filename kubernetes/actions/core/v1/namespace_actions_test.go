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

func initNamespaces() *actionsCoreV1.Namespace {
	info := map[string]string{
		"1": "default",
		"2": "beta",
		"3": "tools",
		"4": "istio-system",
	}
	objects := []runtime.Object{}
	objectsDynamic := []runtime.Object{}
	for _, name := range info {
		namespace := corev1.NewNamespaceBuilder(name)
		buildedNamespace := namespace.SetAnnotations(map[string]string{}).Build()
		objects = append(objects, buildedNamespace)
	}
	client := fake.NewSimpleClientset(objects...)
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).CoreV1.Namespace
	return actions
}

func TestGetNamespace(t *testing.T) {
	actions := initNamespaces()
	namespace, _ := actions.Get("tools")
	assert.Equal(t, "tools", namespace.Name)
}

func TestCreateNamespace(t *testing.T) {
	actions := initNamespaces()
	builder := corev1.NewNamespaceBuilder("andres")
	actions.Create(builder.Build())
	namespacesList, _ := actions.List()
	assert.Equal(t, 5, len(namespacesList.Items))
	assert.Equal(t, "andres", namespacesList.Items[0].Name)
}

func TestUpdateNamespace(t *testing.T) {
	actions := initNamespaces()
	namespace, _ := actions.Get("tools")
	namespace.Labels = map[string]string{"use": "tools"}
	actions.Update(namespace)
	namespaceUpdated, _ := actions.Get("tools")
	assert.Equal(t, map[string]string{"use": "tools"}, namespaceUpdated.Labels)
}

func TestDeleteNamespace(t *testing.T) {
	actions := initNamespaces()
	actions.Delete("beta")
	namespacesList, _ := actions.List()
	assert.Equal(t, 3, len(namespacesList.Items))
	for _, namespace := range namespacesList.Items {
		assert.NotEqual(t, "beta", namespace.Name)
	}
}

func TestListNamespaces(t *testing.T) {
	actions := initNamespaces()
	namespacesList, _ := actions.List()
	assert.Equal(t, 4, len(namespacesList.Items))
}
