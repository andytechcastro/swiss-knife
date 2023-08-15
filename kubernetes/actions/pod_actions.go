package actions

import (
	"context"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreInterface "k8s.io/client-go/kubernetes/typed/core/v1"
)

// Pod struct for actions pod
type Pod struct {
	client           coreInterface.CoreV1Interface
	CurrentNamespace string
}

// NewPodAction return an actions pod
func NewPodAction(client coreInterface.CoreV1Interface) *Pod {
	return &Pod{
		client:           client,
		CurrentNamespace: "",
	}
}

// Namespace set namespace
func (p *Pod) Namespace(namespace string) *Pod {
	p.CurrentNamespace = namespace
	return p
}

// Create Create a pods in the client
func (p *Pod) Create(pod *apiv1.Pod) error {
	if pod == nil {
		return errorPodEmpty
	}
	_, err := p.client.Pods(p.CurrentNamespace).Create(
		context.TODO(),
		pod,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// Update Update a pods in the client
func (p *Pod) Update(pod *apiv1.Pod) error {
	if pod == nil {
		return errorPodEmpty
	}
	_, err := p.client.Pods(p.CurrentNamespace).Update(
		context.TODO(),
		pod,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete a pod in the client
func (p *Pod) Delete(podName string) error {
	if podName == "" {
		return errorNameEmpty
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := p.client.Pods(p.CurrentNamespace).Delete(
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

// Get Get a pod from the client
func (p *Pod) Get(podName string) (*apiv1.Pod, error) {
	if podName == "" {
		return nil, errorNameEmpty
	}
	pod, err := p.client.Pods(p.CurrentNamespace).Get(
		context.TODO(),
		podName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}

	return pod, nil
}

// List List all pods in a namespace
func (p *Pod) List() (*apiv1.PodList, error) {
	podList, err := p.client.Pods(p.CurrentNamespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return podList, nil
}
