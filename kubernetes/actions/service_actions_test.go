package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

func initServices() *actions.Service {
	info := map[string]string{
		"service1": "default",
		"service2": "default",
		"service3": "default",
		"service4": "default",
	}
	objects := []runtime.Object{}
	for name, namespace := range info {
		service := corev1.NewServiceBuilder(name)
		ports := builders.Ports{
			Protocol:   "TCP",
			Port:       80,
			TargetPort: intstr.FromInt(8080),
		}
		buildedService, _ := service.SetNamespace(namespace).
			SetSelector(map[string]string{"service": name}).
			AddPorts(ports.Build()).
			Build()

		objects = append(objects, buildedService)
	}
	client := fake.NewSimpleClientset(objects...)
	objectsDynamic := []runtime.Object{}

	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).Service
	actions.Namespace("default")
	return actions
}

func TestGetService(t *testing.T) {
	actions := initServices()
	service, _ := actions.Get("service1")
	assert.Equal(t, "service1", service.Name)
}

func TestUpdateService(t *testing.T) {
	actions := initServices()
	service, _ := actions.Get("service1")
	service.Spec.Ports[0].TargetPort = intstr.FromInt(8081)
	service.Spec.Ports[0].Port = 81
	actions.Update(service)
	serviceUpdated, _ := actions.Get("service1")
	assert.Equal(t, intstr.FromInt(8081), serviceUpdated.Spec.Ports[0].TargetPort)
	assert.Equal(t, int32(81), serviceUpdated.Spec.Ports[0].Port)
}

func TestCreateService(t *testing.T) {
	actions := initServices()
	service := corev1.NewServiceBuilder("service5")
	ports := builders.Ports{
		Protocol:   "TCP",
		Port:       80,
		TargetPort: intstr.FromInt(8080),
	}
	buildedService, _ := service.SetNamespace("default").
		SetSelector(map[string]string{"service": "service5"}).
		AddPorts(ports.Build()).
		Build()
	actions.Create(buildedService)
	newService, _ := actions.Get("service5")
	services, _ := actions.List()
	assert.Equal(t, "service5", newService.Name)
	assert.Equal(t, 5, len(services.Items))
}

func TestCreateServiceFailed(t *testing.T) {
	actions := initServices()
	service := corev1.NewServiceBuilder("service5")
	ports := builders.Ports{
		Protocol:   "TCP",
		Port:       80,
		TargetPort: intstr.FromInt(8080),
	}
	buildedService, _ := service.SetNamespace("beta").
		SetSelector(map[string]string{"service": "service5"}).
		AddPorts(ports.Build()).
		Build()
	err := actions.Create(buildedService)
	assert.NotNil(t, err)
}

func TestDeleteService(t *testing.T) {
	actions := initServices()
	actions.Delete("service3")
	services, _ := actions.List()
	assert.Equal(t, 3, len(services.Items))
	for _, service := range services.Items {
		assert.NotEqual(t, "service3", service.Name)
	}
}

func TestListService(t *testing.T) {
	actions := initServices()
	services, _ := actions.List()
	assert.Equal(t, 4, len(services.Items))
}
