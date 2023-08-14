package actions

import (
	"context"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreInterface "k8s.io/client-go/kubernetes/typed/core/v1"
)

// ServiceAccount for manager namespace actions
type ServiceAccount struct {
	client coreInterface.ServiceAccountInterface
}

// NewServiceAccountAction get ServiceAccount action
func NewServiceAccountAction(client coreInterface.ServiceAccountInterface) *ServiceAccount {
	return &ServiceAccount{
		client: client,
	}
}

// Get get namespace
func (sa *ServiceAccount) Get(name string) (*apiv1.ServiceAccount, error) {
	serviceAccount, err := sa.client.Get(
		context.TODO(),
		name,
		metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return serviceAccount, nil
}

// Create Create an ServiceAccount
func (sa *ServiceAccount) Create(serviceAccount *apiv1.ServiceAccount) error {
	_, err := sa.client.Create(
		context.TODO(),
		serviceAccount,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Update Update a pods in the client
func (sa *ServiceAccount) Update(serviceAccount *apiv1.ServiceAccount) error {
	if serviceAccount == nil {
		return errorPodEmpty
	}
	_, err := sa.client.Update(
		context.TODO(),
		serviceAccount,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete a pod in the client
func (sa *ServiceAccount) Delete(saName string) error {
	if saName == "" {
		return errorNameEmpty
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := sa.client.Delete(
		context.TODO(),
		saName,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// List List all pods in a namespace
func (sa *ServiceAccount) List() (*apiv1.ServiceAccountList, error) {
	serviceAccountList, err := sa.client.List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return serviceAccountList, nil
}
