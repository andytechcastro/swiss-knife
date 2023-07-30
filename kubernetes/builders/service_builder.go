package builders

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// Service struct for service kubernetes resource
type Service struct {
	Name         string
	Namespace    string
	Type         apiv1.ServiceType
	Selector     map[string]string
	ClusterIP    string
	Labels       map[string]string
	Annotations  map[string]string
	Ports        []apiv1.ServicePort
	ExternalName string
	Service      *apiv1.Service
}

// NewServiceBuilder return a service
func NewServiceBuilder() *Service {
	return &Service{
		Type: apiv1.ServiceTypeClusterIP,
	}
}

// SetName Set the name of the service
func (s *Service) SetName(name string) {
	s.Name = name
}

// SetNamespace Set the namespace of the service
func (s *Service) SetNamespace(namespace string) {
	s.Namespace = namespace
}

// SetType Set the type of a service
// The options could be "ClusterIP", "NodePort", "LoadBalancer" or "ExternalName"
// or use this contants https://pkg.go.dev/k8s.io/api/core/v1#ServiceType
func (s *Service) SetType(t apiv1.ServiceType) {
	s.Type = t
}

// SetSelector Set the selector of a service
func (s *Service) SetSelector(selector map[string]string) {
	s.Selector = selector
}

// SetClusterIP Set the clusterIP of a service
func (s *Service) SetClusterIP(clusterIP string) {
	s.ClusterIP = clusterIP
}

// SetLabels Set labels for the service
func (s *Service) SetLabels(labels map[string]string) {
	s.Labels = labels
}

// AddPorts add ports to the service
func (s *Service) AddPorts(ports *apiv1.ServicePort) {
	s.Ports = append(s.Ports, *ports)
}

// SetAnnotations Set annotations for the service
func (s *Service) SetAnnotations(annotations map[string]string) {
	s.Annotations = annotations
}

// Build Build a service with the data
func (s *Service) Build() (*apiv1.Service, error) {
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	service := &apiv1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        s.Name,
			Namespace:   s.Namespace,
			Labels:      s.Labels,
			Annotations: s.Annotations,
		},
		Spec: apiv1.ServiceSpec{
			Type:         s.Type,
			Selector:     s.Selector,
			ExternalName: s.ExternalName,
			Ports:        s.Ports,
		},
	}
	s.Service = service
	return service, nil
}

// ToYaml Trasnform the struct in yaml
func (s *Service) ToYaml() []byte {
	yaml, err := yaml.Marshal(s.Service)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}

// Validate validate the values for build a service
func (s *Service) Validate() error {
	if s.Type == apiv1.ServiceTypeClusterIP || s.Type == apiv1.ServiceTypeNodePort {
		if s.Selector == nil {
			return errorSelectorEmpty
		} else if s.Ports == nil {
			return errorPortsEmpty
		}
	} else if s.Type == apiv1.ServiceTypeExternalName {
		if s.ExternalName == "" {
			return errorExternalNameEmpty
		}
	}
	return nil
}
