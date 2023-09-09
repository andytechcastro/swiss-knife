package actions

import (
	actionsAppsV1 "github.com/andytechcastro/swiss-knife/kubernetes/actions/apps/v1"
	actionsBatchV1 "github.com/andytechcastro/swiss-knife/kubernetes/actions/batch/v1"
	actionsCoreV1 "github.com/andytechcastro/swiss-knife/kubernetes/actions/core/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Actions kubernetes actions
type Actions struct {
	client  kubernetes.Interface
	config  *rest.Config
	Custom  *Custom
	CoreV1  *actionsCoreV1.CoreV1
	BatchV1 *actionsBatchV1.BatchV1
	AppsV1  *actionsAppsV1.AppsV1
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
	batchV1Client := clientSet.BatchV1()
	return &Actions{
		client:  clientSet,
		config:  config,
		Custom:  NewCustomActions(dynamicClient),
		CoreV1:  actionsCoreV1.NewCoreV1(coreV1Client),
		BatchV1: actionsBatchV1.NewBatchV1(batchV1Client),
		AppsV1:  actionsAppsV1.NewAppsV1(appsV1Client),
	}
}
