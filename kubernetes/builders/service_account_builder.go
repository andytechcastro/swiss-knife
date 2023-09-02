package builders

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// ServiceAccount struct for manage ServiceAccount
type ServiceAccount struct {
	Name           string
	Namespace      string
	Labels         map[string]string
	Annotations    map[string]string
	ServiceAccount *corev1.ServiceAccount
}

// NewServiceAccountBuilder return ServiceAccountBuilder
func NewServiceAccountBuilder() *ServiceAccount {
	return &ServiceAccount{}
}

// SetName set name for service account
func (sa *ServiceAccount) SetName(name string) *ServiceAccount {
	sa.Name = name
	return sa
}

// SetNamespace set namespace for service account
func (sa *ServiceAccount) SetNamespace(namespace string) *ServiceAccount {
	sa.Namespace = namespace
	return sa
}

// SetLabels set labels for service account
func (sa *ServiceAccount) SetLabels(labels map[string]string) *ServiceAccount {
	sa.Labels = labels
	return sa
}

// SetAnnotations set annotations for service account
func (sa *ServiceAccount) SetAnnotations(annotations map[string]string) *ServiceAccount {
	sa.Annotations = annotations
	return sa
}

// Build Build Namespace object
func (sa *ServiceAccount) Build() *corev1.ServiceAccount {
	serviceAccount := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:        sa.Name,
			Labels:      sa.Labels,
			Annotations: sa.Annotations,
			Namespace:   sa.Namespace,
		},
	}
	sa.ServiceAccount = serviceAccount
	return sa.ServiceAccount
}

// ToYaml Convert the object in yaml
func (sa *ServiceAccount) ToYaml() []byte {
	yaml, err := yaml.Marshal(sa.ServiceAccount)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
