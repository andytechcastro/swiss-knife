package actions

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Actions kubernetes actions
type Actions struct {
	client         kubernetes.Interface
	config         *rest.Config
	Namespace      *Namespace
	ServiceAccount *ServiceAccount
	Service        *Service
	Pod            *Pod
	Deployment     *Deployment
	ReplicaSet     *ReplicaSet
	Custom         *Custom
	ConfigMap      *ConfigMap
	Secret         *Secret
	Job            *Job
	CronJob        *CronJob
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
		client:         clientSet,
		config:         config,
		Namespace:      NewNamespaceAction(coreV1Client.Namespaces()),
		Service:        NewServiceAction(coreV1Client),
		Deployment:     NewDeploymentAction(appsV1Client),
		ReplicaSet:     NewReplicaSetAction(appsV1Client),
		Pod:            NewPodAction(coreV1Client),
		ServiceAccount: NewServiceAccountAction(coreV1Client),
		Custom:         NewCustomActions(dynamicClient),
		ConfigMap:      NewConfigMapAction(coreV1Client),
		Secret:         NewSecretAction(coreV1Client),
		Job:            NewJobAction(batchV1Client),
		CronJob:        NewCronJobAction(batchV1Client),
	}
}
