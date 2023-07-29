package actions

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateDeployment Create a deployment in the client
func (a *Actions) CreateDeployment(deployment *appsv1.Deployment) error {
	if deployment == nil {
		return errorDeploymentEmpty
	}
	_, err := a.client.AppsV1().Deployments(a.Namespace).Create(
		context.TODO(),
		deployment,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdateDeployment Update a deployment in the client
func (a *Actions) UpdateDeployment(deployment *appsv1.Deployment) error {
	if deployment == nil {
		return errorDeploymentEmpty
	}
	_, err := a.client.AppsV1().Deployments(a.Namespace).Update(
		context.TODO(),
		deployment,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDeployment Delete a deployment in the client
func (a *Actions) DeleteDeployment(deploymentName string) error {
	if deploymentName == "" {
		return errorNameEmpty
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := a.client.AppsV1().Deployments(a.Namespace).Delete(
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

// GetDeployment Get a deployment from the client
func (a *Actions) GetDeployment(deploymentName string) (*appsv1.Deployment, error) {
	if deploymentName == "" {
		return nil, errorNameEmpty
	}
	deployment, err := a.client.AppsV1().Deployments(a.Namespace).Get(
		context.TODO(),
		deploymentName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

// ListDeployment List all deployments in a namespace
func (a *Actions) ListDeployment() (*appsv1.DeploymentList, error) {
	namespace := a.Namespace
	if a.AllNamespaces {
		namespace = ""
	}
	deploymentList, err := a.client.AppsV1().Deployments(namespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return deploymentList, nil
}
