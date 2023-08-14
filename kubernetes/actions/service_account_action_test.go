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

func initServiceAccounts() *actions.Actions {
	info := map[string]string{
		"service1": "default",
		"service2": "default",
		"service3": "default",
		"service4": "default",
	}
	objects := []runtime.Object{}
	for name, namespace := range info {
		sa := builders.NewServiceAccountBuilder()
		sas := sa.SetName(name).SetNamespace(namespace).Build()
		objects = append(objects, sas)
	}
	client := fake.NewSimpleClientset(objects...)
	objectsDynamic := []runtime.Object{}

	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	actions := actions.NewTestActions(client, dynamicClient, &rest.Config{})

	return actions
}

func TestGetServiceAccounts(t *testing.T) {
	actions := initServiceAccounts()
	serviceAccount, _ := actions.ServiceAccount.Get("service3")
	assert.Equal(t, "service3", serviceAccount.Name)
}

func TestCreateServiceAccount(t *testing.T) {
	actions := initServiceAccounts()
	serviceAccount := builders.NewServiceAccountBuilder()
	buildedServiceAccount := serviceAccount.SetName("service5").
		SetNamespace("default").
		SetLabels(map[string]string{"label": "my-label"}).
		SetAnnotations(map[string]string{"annotation": "my-annotation"}).
		Build()

	actions.ServiceAccount.Create(buildedServiceAccount)
	newServiceAccount, _ := actions.ServiceAccount.Get("service5")
	serviceAccounts, _ := actions.ServiceAccount.List()
	assert.Equal(t, "service5", newServiceAccount.Name)
	assert.Equal(t, 5, len(serviceAccounts.Items))
}

func TestUpdateServiceAccount(t *testing.T) {
	actions := initServiceAccounts()
	serviceAccount, _ := actions.ServiceAccount.Get("service1")
	serviceAccount.Labels = map[string]string{
		"label": "my-label",
	}
	serviceAccount.Annotations = map[string]string{
		"annotation": "my-annotation",
	}
	actions.ServiceAccount.Update(serviceAccount)
	serviceAccountUpdated, _ := actions.ServiceAccount.Get("service1")
	assert.Equal(t, map[string]string{"label": "my-label"}, serviceAccountUpdated.Labels)
	assert.Equal(t, map[string]string{"annotation": "my-annotation"}, serviceAccountUpdated.Annotations)
}

func TestDeleteServiceAccount(t *testing.T) {
	actions := initServiceAccounts()
	actions.ServiceAccount.Delete("service3")
	services, _ := actions.ServiceAccount.List()
	assert.Equal(t, 3, len(services.Items))
	for _, service := range services.Items {
		assert.NotEqual(t, "service3", service.Name)
	}
}

func TestListServiceAccount(t *testing.T) {
	actions := initServiceAccounts()
	serviceAccounts, _ := actions.ServiceAccount.List()
	assert.Equal(t, 4, len(serviceAccounts.Items))
}
