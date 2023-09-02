package builders

import (
	"encoding/json"
)

// Metadata struct for manage metadata
type Metadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// NewMetadata return Metadata struct
func NewMetadata(name string) *Metadata {
	return &Metadata{
		Name: name,
	}
}

// SetNamespace set namespace of metadata
func (m *Metadata) SetNamespace(namespace string) *Metadata {
	m.Namespace = namespace
	return m
}

// SetLabels set labels of metadata
func (m *Metadata) SetLabels(labels map[string]string) *Metadata {
	m.Labels = labels
	return m
}

// SetAnnotations set annotations of metadata
func (m *Metadata) SetAnnotations(annotations map[string]string) *Metadata {
	m.Annotations = annotations
	return m
}

// ToMap Convert Metadata struct to interface
func (m *Metadata) ToMap() map[string]interface{} {
	var toMap map[string]interface{}
	data, _ := json.Marshal(m)
	json.Unmarshal(data, &toMap)
	return toMap
}
