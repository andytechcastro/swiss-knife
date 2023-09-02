package builders

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// ConfigMap struct for configmap kubernetes resource
type ConfigMap struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
	Data        map[string]string
	Immutable   bool
	ConfigMap   *corev1.ConfigMap
}

// NewConfigMapBuilder return a ConfigMap
// Immutable is false for default
func NewConfigMapBuilder(name string) *ConfigMap {
	return &ConfigMap{
		Name: name,
	}
}

// SetNamespace set the configmap namespace
func (c *ConfigMap) SetNamespace(namespace string) *ConfigMap {
	c.Namespace = namespace
	return c
}

// SetLabels set labels in the configmap
func (c *ConfigMap) SetLabels(labels map[string]string) *ConfigMap {
	c.Labels = labels
	return c
}

// SetAnnotations set labels in the configmap
func (c *ConfigMap) SetAnnotations(annotations map[string]string) *ConfigMap {
	c.Annotations = annotations
	return c
}

// SetData set the data of the configmap
func (c *ConfigMap) SetData(data map[string]string) *ConfigMap {
	c.Data = data
	return c
}

// AddData add data to data configmap
func (c *ConfigMap) AddData(key string, value string) *ConfigMap {
	c.Data[key] = value
	return c
}

// SetImmutable set the configmap immutable
func (c *ConfigMap) SetImmutable() *ConfigMap {
	c.Immutable = true
	return c
}

// Build build a configmap
func (c *ConfigMap) Build() *corev1.ConfigMap {
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.Name,
			Namespace: c.Namespace,
		},
		Data:      c.Data,
		Immutable: &c.Immutable,
	}
	c.ConfigMap = configMap
	return configMap
}

// ToYaml convert deployment struct to kubernetes yaml
func (c *ConfigMap) ToYaml() []byte {
	yaml, err := yaml.Marshal(c.ConfigMap)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
