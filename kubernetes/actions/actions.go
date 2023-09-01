package actions

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Actions kubernetes actions
type Actions struct {
	client           kubernetes.Interface
	config           *rest.Config
	CurrentNamespace string
	AllNamespaces    bool
	Namespace        *Namespace
	ServiceAccount   *ServiceAccount
	Service          *Service
	Pod              *Pod
	Deployment       *Deployment
	Custom           *Custom
	ConfigMap        *ConfigMap
	Secret           *Secret
}

// NewActions get an actions interface
func NewActions(config *rest.Config) (*Actions, error) {
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return GetActionFilled(client, dynamicClient, config), nil
}

// GetActionFilled return an action from dynamic client
func GetActionFilled(clientSet kubernetes.Interface, dynamicClient dynamic.Interface, config *rest.Config) *Actions {
	coreV1Client := clientSet.CoreV1()
	appsV1Client := clientSet.AppsV1()
	return &Actions{
		client:           clientSet,
		config:           config,
		CurrentNamespace: "default",
		AllNamespaces:    false,
		Namespace:        NewNamespaceAction(coreV1Client.Namespaces()),
		Service:          NewServiceAction(coreV1Client),
		Deployment:       NewDeploymentAction(appsV1Client),
		Pod:              NewPodAction(coreV1Client),
		ServiceAccount:   NewServiceAccountAction(coreV1Client),
		Custom:           NewCustomActions(dynamicClient),
		ConfigMap:        NewConfigMapAction(coreV1Client),
		Secret:           NewSecretAction(coreV1Client),
	}
}

// EnableAllNamespaces work with all namespaces
func (a *Actions) EnableAllNamespaces() {
	a.AllNamespaces = true
}

// DisableAllNamespaces work with only one namespace
func (a *Actions) DisableAllNamespaces() {
	a.AllNamespaces = false
}

// SetNamespace Set the namespace for use
func (a *Actions) SetNamespace(namespace string) {
	a.CurrentNamespace = namespace
}

// GetCurrentNamespace get the seted namespace
func (a *Actions) GetCurrentNamespace() string {
	return a.CurrentNamespace
}

// GetAllNamespaces Get if AllNamespaces is activated
func (a *Actions) GetAllNamespaces() bool {
	return a.AllNamespaces
}
