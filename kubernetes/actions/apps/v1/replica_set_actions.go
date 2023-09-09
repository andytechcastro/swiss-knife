package actions

import (
	"context"

	"github.com/andytechcastro/swiss-knife/errors"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsInterface "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// ReplicaSet strct for replicaSets action
type ReplicaSet struct {
	client           appsInterface.AppsV1Interface
	CurrentNamespace string
}

// NewReplicaSetAction get a replicaSet action
func NewReplicaSetAction(client appsInterface.AppsV1Interface) *ReplicaSet {
	return &ReplicaSet{
		client:           client,
		CurrentNamespace: "",
	}
}

// Namespace set namespace
func (d *ReplicaSet) Namespace(namespace string) *ReplicaSet {
	d.CurrentNamespace = namespace
	return d
}

// Create Create a replicaSet in the client
func (d *ReplicaSet) Create(replicaSet *appsv1.ReplicaSet) error {
	if replicaSet == nil {
		return errors.GetEmptyError("ReplicaSet")
	}
	_, err := d.client.ReplicaSets(d.CurrentNamespace).Create(
		context.TODO(),
		replicaSet,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// Update Update a replicaSet in the client
func (d *ReplicaSet) Update(replicaSet *appsv1.ReplicaSet) error {
	if replicaSet == nil {
		return errors.GetEmptyError("ReplicaSet")
	}
	_, err := d.client.ReplicaSets(d.CurrentNamespace).Update(
		context.TODO(),
		replicaSet,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete a replicaSet in the client
func (d *ReplicaSet) Delete(replicaSetName string) error {
	if replicaSetName == "" {
		return errors.GetEmptyError("Name")
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := d.client.ReplicaSets(d.CurrentNamespace).Delete(
		context.TODO(),
		replicaSetName,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Get Get a replicaSet from the client
func (d *ReplicaSet) Get(replicaSetName string) (*appsv1.ReplicaSet, error) {
	if replicaSetName == "" {
		return nil, errors.GetEmptyError("Name")
	}
	replicaSet, err := d.client.ReplicaSets(d.CurrentNamespace).Get(
		context.TODO(),
		replicaSetName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}

	return replicaSet, nil
}

// List List all replicaSets in a namespace
func (d *ReplicaSet) List() (*appsv1.ReplicaSetList, error) {
	replicaSetList, err := d.client.ReplicaSets(d.CurrentNamespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return replicaSetList, nil
}
