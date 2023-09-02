package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/fake"

	"k8s.io/client-go/rest"
)

func initNamespaces() *actions.Namespace {
	client := fake.NewSimpleClientset(&corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "default",
			Annotations: map[string]string{},
		},
	}, &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "beta",
			Annotations: map[string]string{},
		},
	}, &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "tools",
			Annotations: map[string]string{},
		},
	}, &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "istio-system",
			Annotations: map[string]string{},
		},
	})
	objects := []runtime.Object{}
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objects...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).Namespace
	return actions
}

func TestGetNamespace(t *testing.T) {
	actions := initNamespaces()
	namespace, _ := actions.Get("tools")
	assert.Equal(t, "tools", namespace.Name)
}

func TestCreateNamespace(t *testing.T) {
	actions := initNamespaces()
	builder := builders.NewNamespaceBuilder("andres")
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
