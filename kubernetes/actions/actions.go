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
}

// NewActions get an actions interface
func NewActions(client kubernetes.Interface) *Actions {
	return &Actions{
		client:           client,
		CurrentNamespace: "default",
		AllNamespaces:    false,
	}
}

// NewTestActions return an action from dynamic client
func NewTestActions(clientSet kubernetes.Interface, dynamicClient dynamic.Interface, config *rest.Config) *Actions {
	//serviceAccountRes := schema.GroupVersionResource{
	//	Group:    "",
	//	Version:  "v1",
	//	Resource: "serviceaccounts",
	//}
	//serviceAccount := dynamicClient.Resource(serviceAccountRes).Namespace("default")
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
