package builders

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// Pod is the base for work with Pods
type Pod struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
	Containers  []corev1.Container
	Pod         *corev1.Pod
	PodTemplate *corev1.PodTemplateSpec
}

// NewPodBuilder return a pod struct
func NewPodBuilder(name string) *Pod {
	return &Pod{
		Name: name,
	}
}

// SetNamespace Set namespace for deployment
func (p *Pod) SetNamespace(namespace string) *Pod {
	p.Namespace = namespace
	return p
}

// AddContainer Add new container to deployment
func (p *Pod) AddContainer(container corev1.Container) *Pod {
	p.Containers = append(p.Containers, container)
	return p
}

// SetLabels Set labels for deployment
func (p *Pod) SetLabels(labels map[string]string) *Pod {
	p.Labels = labels
	return p
}

// SetAnnotations Set Annotations for deployment
func (p *Pod) SetAnnotations(annotations map[string]string) *Pod {
	p.Annotations = annotations
	return p
}

// Build Build de deployment interface
func (p *Pod) Build() *corev1.Pod {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        p.Name,
			Namespace:   p.Namespace,
			Labels:      p.Labels,
			Annotations: p.Annotations,
		},
		Spec: corev1.PodSpec{
			Containers: p.Containers,
		},
	}
	p.Pod = pod
	return p.Pod
}

// BuildTemplate build pod template
func (p *Pod) BuildTemplate() *corev1.PodTemplateSpec {
	pod := &corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      p.Labels,
			Annotations: p.Annotations,
		},
		Spec: corev1.PodSpec{
			Containers: p.Containers,
		},
	}
	p.PodTemplate = pod
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
