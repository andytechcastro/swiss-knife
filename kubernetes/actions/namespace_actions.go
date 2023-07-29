package actions

import (
	"context"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateNamespace Create a namespace in the client
func (a *Actions) CreateNamespace(namespace *apiv1.Namespace) error {
	if namespace == nil {
		return errorNamespaceEmpty
	}
	_, err := a.client.CoreV1().Namespaces().Create(
		context.TODO(),
		namespace,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// UpdateNamespace Create a namespace in the client
func (a *Actions) UpdateNamespace(namespace *apiv1.Namespace) error {
	if namespace == nil {
		return errorNamespaceEmpty
	}
	_, err := a.client.CoreV1().Namespaces().Update(
		context.TODO(),
		namespace,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteNamespace Delete the namespace in the client
func (a *Actions) DeleteNamespace(namespaceName string) error {
	if namespaceName == "" {
		return errorNameEmpty
	}
	err := a.client.CoreV1().Namespaces().Delete(
		context.TODO(),
		namespaceName,
		metav1.DeleteOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// GetNamespace Get a namespace in the client
func (a *Actions) GetNamespace(namespaceName string) (*apiv1.Namespace, error) {
	if namespaceName == "" {
		return nil, errorNameEmpty
	}
	ns, err := a.client.CoreV1().Namespaces().Get(
		context.TODO(),
		namespaceName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// ListNamespace List the namespaces of a client
func (a *Actions) ListNamespace() (*apiv1.NamespaceList, error) {
	namespaceList, err := a.client.CoreV1().Namespaces().List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return namespaceList, nil
}
