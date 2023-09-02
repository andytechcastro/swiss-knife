package builders

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// Service struct for service kubernetes resource
type Service struct {
	Name         string
	Namespace    string
	Type         corev1.ServiceType
	Selector     map[string]string
	ClusterIP    string
	Labels       map[string]string
	Annotations  map[string]string
	Ports        []corev1.ServicePort
	ExternalName string
	Service      *corev1.Service
}

// NewServiceBuilder return a service
func NewServiceBuilder() *Service {
	return &Service{
		Type: corev1.ServiceTypeClusterIP,
	}
}

// SetName Set the name of the service
func (s *Service) SetName(name string) *Service {
	s.Name = name
	return s
}

// SetNamespace Set the namespace of the service
func (s *Service) SetNamespace(namespace string) *Service {
	s.Namespace = namespace
	return s
}

// SetType Set the type of a service
// The options could be "ClusterIP", "NodePort", "LoadBalancer" or "ExternalName"
// or use this contants https://pkg.go.dev/k8s.io/api/core/v1#ServiceType
func (s *Service) SetType(t corev1.ServiceType) *Service {
	s.Type = t
	return s
}

// SetExternalName set the external name (optional)
func (s *Service) SetExternalName(externalName string) *Service {
	s.ExternalName = externalName
	return s
}

// SetSelector Set the selector of a service
func (s *Service) SetSelector(selector map[string]string) *Service {
	s.Selector = selector
	return s
}

// SetClusterIP Set the clusterIP of a service
func (s *Service) SetClusterIP(clusterIP string) *Service {
	s.ClusterIP = clusterIP
	return s
}

// SetLabels Set labels for the service
func (s *Service) SetLabels(labels map[string]string) *Service {
	s.Labels = labels
	return s
}

// AddPorts add ports to the service
func (s *Service) AddPorts(ports *corev1.ServicePort) *Service {
	s.Ports = append(s.Ports, *ports)
	return s
}

// SetAnnotations Set annotations for the service
func (s *Service) SetAnnotations(annotations map[string]string) *Service {
	s.Annotations = annotations
	return s
}

// Build Build a service with the data
func (s *Service) Build() (*corev1.Service, error) {
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        s.Name,
			Namespace:   s.Namespace,
			Labels:      s.Labels,
			Annotations: s.Annotations,
		},
		Spec: corev1.ServiceSpec{
			Type:         s.Type,
			Selector:     s.Selector,
			ExternalName: s.ExternalName,
			Ports:        s.Ports,
		},
	}
	s.Service = service
	return s.Service, nil
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
	if s.Type == corev1.ServiceTypeClusterIP || s.Type == corev1.ServiceTypeNodePort {
		if s.Selector == nil {
			return errorSelectorEmpty
		} else if s.Ports == nil {
			return errorPortsEmpty
		}
	} else if s.Type == corev1.ServiceTypeExternalName {
		if s.ExternalName == "" {
			return errorExternalNameEmpty
		}
	}
	return nil
}
