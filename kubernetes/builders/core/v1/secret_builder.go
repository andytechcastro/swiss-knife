package builders

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// Secret struct for secret
type Secret struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
	Data        map[string][]byte
	StringData  map[string]string
	Immutable   bool
	Type        corev1.SecretType
	Secret      *corev1.Secret
}

// NewSecretBuilder get a sercret builder
func NewSecretBuilder(name string) *Secret {
	return &Secret{
		Name: name,
	}
}

// SetNamespace set the namespace of the secret
func (s *Secret) SetNamespace(namespace string) *Secret {
	s.Namespace = namespace
	return s
}

// SetLabels set labels in the configmap
func (s *Secret) SetLabels(labels map[string]string) *Secret {
	s.Labels = labels
	return s
}

// SetAnnotations set labels in the configmap
func (s *Secret) SetAnnotations(annotations map[string]string) *Secret {
	s.Annotations = annotations
	return s
}

// SetImmutable set to immutable the secret
func (s *Secret) SetImmutable() *Secret {
	s.Immutable = true
	return s
}

// SetData set the data in a secret
func (s *Secret) SetData(data map[string][]byte) *Secret {
	s.Data = data
	return s
}

// SetStringData set data in the secret
func (s *Secret) SetStringData(data map[string]string) *Secret {
	s.StringData = data
	return s
}

// SetType set the secret type
func (s *Secret) SetType(secretType corev1.SecretType) *Secret {
	s.Type = secretType
	return s
}

// Build a secret
func (s *Secret) Build() *corev1.Secret {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:        s.Name,
			Namespace:   s.Namespace,
			Labels:      s.Labels,
			Annotations: s.Annotations,
		},
		Data:       s.Data,
		StringData: s.StringData,
		Immutable:  &s.Immutable,
		Type:       s.Type,
	}

	s.Secret = secret
	return secret
}

// ToYaml convert secret struct to kubernetes yaml
func (s *Secret) ToYaml() []byte {
	yaml, err := yaml.Marshal(s.Secret)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
