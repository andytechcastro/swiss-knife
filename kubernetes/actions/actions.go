package actions

import "k8s.io/client-go/kubernetes"

// Actions kubernetes actions
type Actions struct {
	client        kubernetes.Interface
	Namespace     string
	AllNamespaces bool
}

// NewActions get an actions interface
func NewActions(client kubernetes.Interface) *Actions {
	return &Actions{
		client:        client,
		Namespace:     "default",
		AllNamespaces: false,
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
	a.Namespace = namespace
}

// GetCurrentNamespace get the seted namespace
func (a *Actions) GetCurrentNamespace() string {
	return a.Namespace
}

// GetAllNamespaces Get if AllNamespaces is activated
func (a *Actions) GetAllNamespaces() bool {
	return a.AllNamespaces
}
