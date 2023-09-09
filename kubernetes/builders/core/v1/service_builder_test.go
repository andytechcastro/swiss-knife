package builders_test

import (
	"testing"

	builderCorev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/yaml"
)

const ServiceName = "my-service"

func initService() *builderCorev1.Service {
	service := builderCorev1.NewServiceBuilder(ServiceName)
	return service
}

func TestBuildServiceClusterIP(t *testing.T) {
	service := initService()
	service.Selector = map[string]string{"service": ServiceName}
	ports := builderCorev1.NewServicePort()
	ports.SetProtocol("TCP").
		SetPort(80).
		SetTargetPort(8080)
	service.AddPorts(ports.Build())
	buildedService, _ := service.Build()
	assert.Equal(t, ServiceName, buildedService.Name)
	assert.Equal(t, corev1.ServiceTypeClusterIP, buildedService.Spec.Type)
	assert.Equal(t, corev1.ProtocolTCP, buildedService.Spec.Ports[0].Protocol)
	assert.Equal(t, int32(80), buildedService.Spec.Ports[0].Port)
	assert.Equal(t, intstr.FromInt(8080), buildedService.Spec.Ports[0].TargetPort)
}

func TestBuildServiceExternalName(t *testing.T) {
	service := initService()
	buildedService, _ := service.SetExternalName(ServiceName).
		SetType("ExternalName").
		Build()
	assert.Equal(t, ServiceName, buildedService.Name)
	assert.Equal(t, corev1.ServiceTypeExternalName, buildedService.Spec.Type)
	assert.Equal(t, ServiceName, buildedService.Spec.ExternalName)
}

func TestBuildServiceExternalNameEmpty(t *testing.T) {
	service := initService()
	service.Type = "ExternalName"
	_, err := service.Build()
	assert.NotNil(t, err)
}

func TestServiceEmptyValue(t *testing.T) {
	service := initService()
	_, err := service.Build()
	assert.NotNil(t, err)
}

func TestServiceToYaml(t *testing.T) {
	service := initService()
	ports := builderCorev1.Ports{
		Protocol:   "TCP",
		Port:       80,
		TargetPort: intstr.FromInt(8080),
	}
	service.SetSelector(map[string]string{"service": ServiceName}).
		AddPorts(ports.Build()).
		Build()
	yamlService := service.ToYaml()
	interfaceResult := map[string]interface{}(
		map[string]interface{}{
			"metadata": map[string]interface{}{
				"creationTimestamp": interface{}(nil),
				"name":              ServiceName,
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{map[string]interface{}{
					"port":       80,
					"protocol":   "TCP",
					"targetPort": 8080,
				}},
				"type": "ClusterIP",
				"selector": map[string]interface{}{
					"service": ServiceName,
				},
			},
			"status": map[string]interface{}{
				"loadBalancer": map[string]interface{}{},
			},
		},
	)
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlResult), string(yamlService))
}
