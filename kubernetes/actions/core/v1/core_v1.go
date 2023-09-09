package actions

import (
	coreInterface "k8s.io/client-go/kubernetes/typed/core/v1"
)

// CoreV1 struct for access to batchv1 api
type CoreV1 struct {
	Namespace      *Namespace
	ServiceAccount *ServiceAccount
	ConfigMap      *ConfigMap
	Secret         *Secret
	Pod            *Pod
	Service        *Service
}

// NewCoreV1 return a batch v1 api
func NewCoreV1(client coreInterface.CoreV1Interface) *CoreV1 {
	return &CoreV1{
		Namespace:      NewNamespaceAction(client.Namespaces()),
		ServiceAccount: NewServiceAccountAction(client),
		ConfigMap:      NewConfigMapAction(client),
		Secret:         NewSecretAction(client),
		Pod:            NewPodAction(client),
		Service:        NewServiceAction(client),
	}
}
