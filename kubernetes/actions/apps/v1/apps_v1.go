package actions

import (
	appsInterface "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// AppsV1 struct for access to batchv1 api
type AppsV1 struct {
	Deployment *Deployment
	ReplicaSet *ReplicaSet
}

// NewAppsV1 return a batch v1 api
func NewAppsV1(client appsInterface.AppsV1Interface) *AppsV1 {
	return &AppsV1{
		Deployment: NewDeploymentAction(client),
		ReplicaSet: NewReplicaSetAction(client),
	}
}
