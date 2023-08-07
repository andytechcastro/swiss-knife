package builders

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// Namespace is the struct for namespace
type Namespace struct {
	Name        string
	Labels      map[string]string
	Annotations map[string]string
	Namespace   *apiv1.Namespace
}

// NewNamespaceBuilder return a namespace builder
func NewNamespaceBuilder() *Namespace {
	return &Namespace{}
}

// SetName set name of a namespace
func (n *Namespace) SetName(name string) *Namespace {
	n.Name = name
	return n
}

// SetLabels set labels of a namespace
func (n *Namespace) SetLabels(labels map[string]string) *Namespace {
	n.Labels = labels
	return n
}

// SetAnnotations sets annotations of a namespace
func (n *Namespace) SetAnnotations(annotations map[string]string) *Namespace {
	n.Annotations = annotations
	return n
}

// Build build a namespace
func (n *Namespace) Build() *apiv1.Namespace {
	namespace := &apiv1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        n.Name,
			Labels:      n.Labels,
			Annotations: n.Annotations,
		},
	}
	n.Namespace = namespace
	return n.Namespace
}

// ToYaml conver namespace struct to kubernetes yaml
func (n *Namespace) ToYaml() []byte {
	yaml, err := yaml.Marshal(n.Namespace)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
