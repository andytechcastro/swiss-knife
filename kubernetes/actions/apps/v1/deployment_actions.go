package actions

import (
	"context"

	"github.com/andytechcastro/swiss-knife/errors"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsInterface "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// Deployment strct for deployments action
type Deployment struct {
	client           appsInterface.AppsV1Interface
	CurrentNamespace string
}

// NewDeploymentAction get a deployment action
func NewDeploymentAction(client appsInterface.AppsV1Interface) *Deployment {
	return &Deployment{
		client:           client,
		CurrentNamespace: "",
	}
}

// Namespace set namespace
func (d *Deployment) Namespace(namespace string) *Deployment {
	d.CurrentNamespace = namespace
	return d
}

// Create Create a deployment in the client
func (d *Deployment) Create(deployment *appsv1.Deployment) error {
	if deployment == nil {
		return errors.GetEmptyError("Deployment")
	}
	_, err := d.client.Deployments(d.CurrentNamespace).Create(
		context.TODO(),
		deployment,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// Update Update a deployment in the client
func (d *Deployment) Update(deployment *appsv1.Deployment) error {
	if deployment == nil {
		return errors.GetEmptyError("Deployment")
	}
	_, err := d.client.Deployments(d.CurrentNamespace).Update(
		context.TODO(),
		deployment,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete a deployment in the client
func (d *Deployment) Delete(deploymentName string) error {
	if deploymentName == "" {
		return errors.GetEmptyError("Name")
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := d.client.Deployments(d.CurrentNamespace).Delete(
		context.TODO(),
		deploymentName,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Get Get a deployment from the client
func (d *Deployment) Get(deploymentName string) (*appsv1.Deployment, error) {
	if deploymentName == "" {
		return nil, errors.GetEmptyError("Name")
	}
	deployment, err := d.client.Deployments(d.CurrentNamespace).Get(
		context.TODO(),
		deploymentName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

// List List all deployments in a namespace
func (d *Deployment) List() (*appsv1.DeploymentList, error) {
	deploymentList, err := d.client.Deployments(d.CurrentNamespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return deploymentList, nil
}
