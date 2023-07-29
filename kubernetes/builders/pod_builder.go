package builders

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// Pod is the base for work with Pods
type Pod struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
	Containers  []apiv1.Container
	Pod         *apiv1.Pod
}

// NewPodBuilder return a pod struct
func NewPodBuilder() *Pod {
	return &Pod{}
}

// SetName Set name for deployment
func (p *Pod) SetName(name string) {
	p.Name = name
}

// SetNamespace Set namespace for deployment
func (p *Pod) SetNamespace(namespace string) {
	p.Namespace = namespace
}

// AddContainer Add new container to deployment
func (p *Pod) AddContainer(container apiv1.Container) {
	p.Containers = append(p.Containers, container)
}

// SetLabels Set labels for deployment
func (p *Pod) SetLabels(labels map[string]string) {
	p.Labels = labels
}

// SetAnnotations Set Annotations for deployment
func (p *Pod) SetAnnotations(annotations map[string]string) {
	p.Annotations = annotations
}

// Build Build de deployment interface
func (p *Pod) Build() *apiv1.Pod {
	pod := &apiv1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        p.Name,
			Namespace:   p.Namespace,
			Labels:      p.Labels,
			Annotations: p.Annotations,
		},
		Spec: apiv1.PodSpec{
			Containers: p.Containers,
		},
	}
	p.Pod = pod
	return pod
}

// ToYaml convert deployment struct to kubernetes yaml
func (p *Pod) ToYaml() []byte {
	yaml, err := yaml.Marshal(p.Pod)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
