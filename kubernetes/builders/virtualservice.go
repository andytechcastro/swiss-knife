package builders

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// VirtualService struct for virtualService
type VirtualService struct {
	Name           string
	Namespace      string
	Host           string
	HTTP           []interface{}
	Gateways       []interface{}
	VirtualService unstructured.Unstructured
}

// SetName set the name of the virtualservice
func (v *VirtualService) SetName(name string) {
	v.Name = name
}

// SetNamespace Set namespace of the virtualservice
func (v *VirtualService) SetNamespace(namespace string) {
	v.Namespace = namespace
}

// SetHost set host for virtualservice
func (v *VirtualService) SetHost(host string) {
	v.Host = host
}

// Build build a VirtualService with the information
func (v *VirtualService) Build() unstructured.Unstructured {
	virtualService := unstructured.Unstructured{}
	virtualService.Object = map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":      v.Name,
			"namespace": v.Namespace,
		},
		"spec": map[string]interface{}{
			"host":     v.Host,
			"http":     v.HTTP,
			"gateways": v.Gateways,
		},
	}
	virtualService.SetGroupVersionKind(
		schema.GroupVersionKind{
			Group:   "networking.istio.io",
			Kind:    "VirtualService",
			Version: "v1beta1",
		},
	)

	v.VirtualService = virtualService
	return virtualService
}
