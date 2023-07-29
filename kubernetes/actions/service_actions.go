package actions

import (
	"context"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateService create a service in the client
func (a *Actions) CreateService(service *apiv1.Service) error {
	if service == nil {
		return errorServiceEmpty
	}
	_, err := a.client.CoreV1().Services(a.Namespace).Create(
		context.TODO(),
		service,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// UpdateService a service in the client
func (a *Actions) UpdateService(service *apiv1.Service) error {
	if service == nil {
		return errorServiceEmpty
	}
	_, err := a.client.CoreV1().Services(a.Namespace).Update(
		context.TODO(),
		service,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteService Delete a service in the client
func (a *Actions) DeleteService(serviceName string) error {
	if serviceName == "" {
		return errorNameEmpty
	}
	err := a.client.CoreV1().Services(a.Namespace).Delete(
		context.TODO(),
		serviceName,
		metav1.DeleteOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// GetService Get a service in the client
func (a *Actions) GetService(serviceName string) (*apiv1.Service, error) {
	if serviceName == "" {
		return nil, errorNameEmpty
	}
	service, err := a.client.CoreV1().Services(a.Namespace).Get(
		context.TODO(),
		serviceName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}

	return service, nil
}

// ListService all services in a namespace
func (a *Actions) ListService() (*apiv1.ServiceList, error) {
	serviceList, err := a.client.CoreV1().Services(a.Namespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return serviceList, nil
}
