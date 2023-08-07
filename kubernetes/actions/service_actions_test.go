package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/fake"
)

func initServices() *actions.Actions {
	info := map[string]string{
		"service1": "default",
		"service2": "default",
		"service3": "default",
		"service4": "default",
	}
	objects := []runtime.Object{}
	for name, namespace := range info {
		service := builders.NewServiceBuilder()
		ports := builders.Ports{
			Protocol:   "TCP",
			Port:       80,
			TargetPort: intstr.FromInt(8080),
		}
		buildedService, _ := service.SetName(name).
			SetNamespace(namespace).
			SetSelector(map[string]string{"service": name}).
			AddPorts(ports.Build()).
			Build()

		objects = append(objects, buildedService)
	}
	client := fake.NewSimpleClientset(objects...)

	actions := actions.NewActions(client)
	return actions
}

func TestGetService(t *testing.T) {
	actions := initServices()
	service, _ := actions.GetService("service1")
	assert.Equal(t, "service1", service.Name)
}

func TestUpdateService(t *testing.T) {
	actions := initServices()
	service, _ := actions.GetService("service1")
	service.Spec.Ports[0].TargetPort = intstr.FromInt(8081)
	service.Spec.Ports[0].Port = 81
	actions.UpdateService(service)
	serviceUpdated, _ := actions.GetService("service1")
	assert.Equal(t, intstr.FromInt(8081), serviceUpdated.Spec.Ports[0].TargetPort)
	assert.Equal(t, int32(81), serviceUpdated.Spec.Ports[0].Port)
}

func TestCreateService(t *testing.T) {
	actions := initServices()
	service := builders.NewServiceBuilder()
	ports := builders.Ports{
		Protocol:   "TCP",
		Port:       80,
		TargetPort: intstr.FromInt(8080),
	}
	buildedService, _ := service.SetName("service5").
		SetNamespace("default").
		SetSelector(map[string]string{"service": "service5"}).
		AddPorts(ports.Build()).
		Build()
	actions.CreateService(buildedService)
	newService, _ := actions.GetService("service5")
	services, _ := actions.ListService()
	assert.Equal(t, "service5", newService.Name)
	assert.Equal(t, 5, len(services.Items))
}

func TestCreateServiceFailed(t *testing.T) {
	actions := initServices()
	service := builders.NewServiceBuilder()
	ports := builders.Ports{
		Protocol:   "TCP",
		Port:       80,
		TargetPort: intstr.FromInt(8080),
	}
	buildedService, _ := service.SetName("service5").
		SetNamespace("beta").
		SetSelector(map[string]string{"service": "service5"}).
		AddPorts(ports.Build()).
		Build()
	err := actions.CreateService(buildedService)
	assert.NotNil(t, err)
}

func TestDeleteService(t *testing.T) {
	actions := initServices()
	actions.DeleteService("service3")
	services, _ := actions.ListService()
	assert.Equal(t, 3, len(services.Items))
	for _, service := range services.Items {
		assert.NotEqual(t, "service3", service.Name)
	}
}

func TestListService(t *testing.T) {
	actions := initServices()
	services, _ := actions.ListService()
	assert.Equal(t, 4, len(services.Items))
}
