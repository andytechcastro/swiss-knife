package actions

import (
	"context"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// Namespace struct for namespace action
type Namespace struct {
	client corev1.NamespaceInterface
}

// NewNamespaceAction return a namespace action
func NewNamespaceAction(client corev1.NamespaceInterface) *Namespace {
	return &Namespace{
		client: client,
	}
}

// Create Create a namespace in the client
func (n *Namespace) Create(namespace *apiv1.Namespace) error {
	if namespace == nil {
		return errorNamespaceEmpty
	}
	_, err := n.client.Create(
		context.TODO(),
		namespace,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Update Update a namespace in the client
func (n *Namespace) Update(namespace *apiv1.Namespace) error {
	if namespace == nil {
		return errorNamespaceEmpty
	}
	_, err := n.client.Update(
		context.TODO(),
		namespace,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete the namespace in the client
func (n *Namespace) Delete(namespaceName string) error {
	if namespaceName == "" {
		return errorNameEmpty
	}
	err := n.client.Delete(
		context.TODO(),
		namespaceName,
		metav1.DeleteOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// Get Get a namespace in the client
func (n *Namespace) Get(namespaceName string) (*apiv1.Namespace, error) {
	if namespaceName == "" {
		return nil, errorNameEmpty
	}
	ns, err := n.client.Get(
		context.TODO(),
		namespaceName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// List List the namespaces of a client
func (n *Namespace) List() (*apiv1.NamespaceList, error) {
	namespaceList, err := n.client.List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return namespaceList, nil
}
