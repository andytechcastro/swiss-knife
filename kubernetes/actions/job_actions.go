package actions

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsInterface "k8s.io/client-go/kubernetes/typed/batch/v1"
)

// Job strct for jobs action
type Job struct {
	client           appsInterface.BatchV1Interface
	CurrentNamespace string
}

// NewJobAction get a job action
func NewJobAction(client appsInterface.BatchV1Interface) *Job {
	return &Job{
		client:           client,
		CurrentNamespace: "",
	}
}

// Namespace set namespace
func (d *Job) Namespace(namespace string) *Job {
	d.CurrentNamespace = namespace
	return d
}

// Create Create a job in the client
func (d *Job) Create(job *batchv1.Job) error {
	if job == nil {
		return errorJobEmpty
	}
	_, err := d.client.Jobs(d.CurrentNamespace).Create(
		context.TODO(),
		job,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// Update Update a job in the client
func (d *Job) Update(job *batchv1.Job) error {
	if job == nil {
		return errorJobEmpty
	}
	_, err := d.client.Jobs(d.CurrentNamespace).Update(
		context.TODO(),
		job,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete a job in the client
func (d *Job) Delete(jobName string) error {
	if jobName == "" {
		return errorNameEmpty
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := d.client.Jobs(d.CurrentNamespace).Delete(
		context.TODO(),
		jobName,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Get Get a job from the client
func (d *Job) Get(jobName string) (*batchv1.Job, error) {
	if jobName == "" {
		return nil, errorNameEmpty
	}
	job, err := d.client.Jobs(d.CurrentNamespace).Get(
		context.TODO(),
		jobName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// List List all jobs in a namespace
func (d *Job) List() (*batchv1.JobList, error) {
	jobList, err := d.client.Jobs(d.CurrentNamespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return jobList, nil
}
