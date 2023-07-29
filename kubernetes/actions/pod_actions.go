package actions

import (
	"context"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreatePod Create a pods in the client
func (a *Actions) CreatePod(pod *apiv1.Pod) error {
	if pod == nil {
		return errorPodEmpty
	}
	_, err := a.client.CoreV1().Pods(a.Namespace).Create(
		context.TODO(),
		pod,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePod Update a pods in the client
func (a *Actions) UpdatePod(pod *apiv1.Pod) error {
	if pod == nil {
		return errorPodEmpty
	}
	_, err := a.client.CoreV1().Pods(a.Namespace).Update(
		context.TODO(),
		pod,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// DeletePod Delete a pod in the client
func (a *Actions) DeletePod(podName string) error {
	if podName == "" {
		return errorNameEmpty
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := a.client.CoreV1().Pods(a.Namespace).Delete(
		context.TODO(),
		podName,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// GetPod Get a pod from the client
func (a *Actions) GetPod(podName string) (*apiv1.Pod, error) {
	if podName == "" {
		return nil, errorNameEmpty
	}
	pod, err := a.client.CoreV1().Pods(a.Namespace).Get(
		context.TODO(),
		podName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}

	return pod, nil
}

// ListPod List all pods in a namespace
func (a *Actions) ListPod() (*apiv1.PodList, error) {
	namespace := a.Namespace
	if a.AllNamespaces {
		namespace = ""
	}
	podList, err := a.client.CoreV1().Pods(namespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return podList, nil
}
