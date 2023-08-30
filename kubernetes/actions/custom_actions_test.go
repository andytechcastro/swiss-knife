package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

func initCustom() *actions.Custom {
	info := map[string]string{
		"service1": "default",
		"service2": "default",
		"service3": "default",
		"service4": "default",
	}
	objects := []runtime.Object{}
	objectsDynamic := []runtime.Object{}
	for name, namespace := range info {
		sas := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "ServiceAccount",
				"metadata": map[string]interface{}{
					"namespace": namespace,
					"name":      name,
				},
			},
		}

		objectsDynamic = append(objectsDynamic, sas)
	}
	client := fake.NewSimpleClientset(objects...)

	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).Custom
	actions.Namespace("default")
	gVR := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "serviceaccounts",
	}
	actions.SetGroupVersionResource(gVR)

	return actions
}

func TestGetCustom(t *testing.T) {
	actions := initCustom()
	custom, _ := actions.Get("service4")
	assert.Equal(t, "service4", custom.Object["metadata"].(map[string]interface{})["name"])
	assert.Equal(t, "default", custom.Object["metadata"].(map[string]interface{})["namespace"])
}

func TestCreateCustom(t *testing.T) {
	actions := initCustom()
	custom := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"namespace": "default",
				"name":      "service5",
			},
		},
	}

	actions.Create(custom)
	newServiceAccount, _ := actions.Get("service5")
	customs, _ := actions.List()
	assert.Equal(t, "service5", newServiceAccount.Object["metadata"].(map[string]interface{})["name"])
	assert.Equal(t, 5, len(customs.Items))
}

func TestUpdateCustom(t *testing.T) {
	actions := initCustom()
	custom, _ := actions.Get("service4")
	custom.SetLabels(map[string]string{
		"label": "my-label",
	})

	actions.Update(custom)
	newCustom, _ := actions.Get("service4")
	assert.Equal(
		t,
		map[string]interface{}{"label": "my-label"},
		newCustom.Object["metadata"].(map[string]interface{})["labels"],
	)
}

func TestDeleteCustom(t *testing.T) {
	actions := initCustom()
	actions.Delete("service2")
	services, _ := actions.List()
	assert.Equal(t, 3, len(services.Items))
	for _, service := range services.Items {
		assert.NotEqual(t, "service2", service.Object["metadata"].(map[string]interface{})["name"])
	}
}

func TestListCustom(t *testing.T) {
	actions := initCustom()
	services, _ := actions.List()
	assert.Equal(t, 4, len(services.Items))
}
