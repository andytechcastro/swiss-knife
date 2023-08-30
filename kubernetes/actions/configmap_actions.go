package actions

import (
	"context"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// ConfigMap struct for configmap actions
type ConfigMap struct {
	client           corev1.CoreV1Interface
	CurrentNamespace string
}

// NewConfigMapAction get an configmap actions
func NewConfigMapAction(client corev1.CoreV1Interface) *ConfigMap {
	return &ConfigMap{
		client:           client,
		CurrentNamespace: "",
	}
}

// Namespace set namespace
func (cm *ConfigMap) Namespace(namespace string) *ConfigMap {
	cm.CurrentNamespace = namespace
	return cm
}

// Get get configmap
func (cm *ConfigMap) Get(name string) (*apiv1.ConfigMap, error) {
	configmap, err := cm.client.ConfigMaps(cm.CurrentNamespace).Get(
		context.TODO(),
		name,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}
	return configmap, nil
}

// Create Create an ConfigMap
func (cm *ConfigMap) Create(configMap *apiv1.ConfigMap) error {
	_, err := cm.client.ConfigMaps(cm.CurrentNamespace).Create(
		context.TODO(),
		configMap,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Update Update a pods in the client
func (cm *ConfigMap) Update(configMap *apiv1.ConfigMap) error {
	if configMap == nil {
		return errorPodEmpty
	}
	_, err := cm.client.ConfigMaps(cm.CurrentNamespace).Update(
		context.TODO(),
		configMap,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete a pod in the client
func (cm *ConfigMap) Delete(cmName string) error {
	if cmName == "" {
		return errorNameEmpty
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := cm.client.ConfigMaps(cm.CurrentNamespace).Delete(
		context.TODO(),
		cmName,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// List List all pods in a namespace
func (cm *ConfigMap) List() (*apiv1.ConfigMapList, error) {
	configMapList, err := cm.client.ConfigMaps(cm.CurrentNamespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return configMapList, nil
}
