package actions

import (
	"context"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// Service is the struct for access to the service actions
type Service struct {
	client           corev1.CoreV1Interface
	CurrentNamespace string
}

// NewServiceAction return a service action
func NewServiceAction(client corev1.CoreV1Interface) *Service {
	return &Service{
		client:           client,
		CurrentNamespace: "",
	}
}

// Namespace set namespace
func (s *Service) Namespace(namespace string) *Service {
	s.CurrentNamespace = namespace
	return s
}

// Create create a service in the client
func (s *Service) Create(service *apiv1.Service) error {
	if service == nil {
		return errorServiceEmpty
	}
	_, err := s.client.Services(s.CurrentNamespace).Create(
		context.TODO(),
		service,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Update a service in the client
func (s *Service) Update(service *apiv1.Service) error {
	if service == nil {
		return errorServiceEmpty
	}
	_, err := s.client.Services(s.CurrentNamespace).Update(
		context.TODO(),
		service,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete a service in the client
func (s *Service) Delete(serviceName string) error {
	if serviceName == "" {
		return errorNameEmpty
	}
	err := s.client.Services(s.CurrentNamespace).Delete(
		context.TODO(),
		serviceName,
		metav1.DeleteOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Get Get a service in the client
func (s *Service) Get(serviceName string) (*apiv1.Service, error) {
	if serviceName == "" {
		return nil, errorNameEmpty
	}
	service, err := s.client.Services(s.CurrentNamespace).Get(
		context.TODO(),
		serviceName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}

	return service, nil
}

// List all services in a namespace
func (s *Service) List() (*apiv1.ServiceList, error) {
	serviceList, err := s.client.Services(s.CurrentNamespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return serviceList, nil
}
