package builders

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// Service struct for service kubernetes resource
type Service struct {
	Name        string
	Namespace   string
	Type        string
	Selector    map[string]string
	ClusterIP   string
	Labels      map[string]string
	Annotations map[string]string
	Service     *apiv1.Service
}

// NewServiceBuilder return a service
func NewServiceBuilder() Service {
	return Service{}
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
func (s *Service) SetType(t string) {
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

// SetAnnotations Set annotations for the service
func (s *Service) SetAnnotations(annotations map[string]string) {
	s.Annotations = annotations
}

// Build Build a service with the data
func (s *Service) Build() *apiv1.Service {
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
			Selector: s.Selector,
		},
	}
	s.Service = service
	return service
}

// ToYaml Trasnform the struct in yaml
func (s *Service) ToYaml() []byte {
	yaml, err := yaml.Marshal(s.Service)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
