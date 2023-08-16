package actions

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

// Custom struct for custom actions
type Custom struct {
	client               dynamic.Interface
	CurrentNamespace     string
	GroupVersionResource schema.GroupVersionResource
}

// NewCustomActions return a custom actions
func NewCustomActions(client dynamic.Interface) *Custom {
	return &Custom{
		client: client,
	}
}

// Namespace set namespace for custom actions
func (c *Custom) Namespace(namespace string) *Custom {
	c.CurrentNamespace = namespace
	return c
}

// SetGroupVersionResource set GroupVersionResource for this actions
func (c *Custom) SetGroupVersionResource(resource schema.GroupVersionResource) *Custom {
	c.GroupVersionResource = resource
	return c
}

// Get get custom resource
func (c *Custom) Get(name string) (*unstructured.Unstructured, error) {
	custom, err := c.client.Resource(c.GroupVersionResource).Namespace(c.CurrentNamespace).Get(
		context.TODO(),
		name,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}
	return custom, nil
}

// Create create a custom resource
func (c *Custom) Create(resource *unstructured.Unstructured) error {
	if resource == nil {
		return errorResourceEmpty
	}
	_, err := c.client.Resource(c.GroupVersionResource).Namespace(c.CurrentNamespace).Create(
		context.TODO(),
		resource,
		metav1.CreateOptions{},
	)
	return err
}

// Update update a custome resource
func (c *Custom) Update(resource *unstructured.Unstructured) error {
	if resource == nil {
		return errorResourceEmpty
	}
	_, err := c.client.Resource(c.GroupVersionResource).Namespace(c.CurrentNamespace).Update(
		context.TODO(),
		resource,
		metav1.UpdateOptions{},
	)
	return err
}

// Delete delete a custom resource
func (c *Custom) Delete(name string) error {
	if name == "" {
		return errorNameEmpty
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := c.client.Resource(c.GroupVersionResource).Namespace(c.CurrentNamespace).Delete(
		context.TODO(),
		name,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		},
	)
	return err
}

// List get list of custom resources
func (c *Custom) List() (*unstructured.UnstructuredList, error) {
	customList, err := c.client.Resource(c.GroupVersionResource).Namespace(c.CurrentNamespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return customList, nil
}
