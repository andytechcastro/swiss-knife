package actions

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsInterface "k8s.io/client-go/kubernetes/typed/batch/v1"
)

// CronJob strct for cronJobs action
type CronJob struct {
	client           appsInterface.BatchV1Interface
	CurrentNamespace string
}

// NewCronJobAction get a cronJob action
func NewCronJobAction(client appsInterface.BatchV1Interface) *CronJob {
	return &CronJob{
		client:           client,
		CurrentNamespace: "",
	}
}

// Namespace set namespace
func (c *CronJob) Namespace(namespace string) *CronJob {
	c.CurrentNamespace = namespace
	return c
}

// Create Create a cronJob in the client
func (c *CronJob) Create(cronJob *batchv1.CronJob) error {
	if cronJob == nil {
		return errorCronJobEmpty
	}
	_, err := c.client.CronJobs(c.CurrentNamespace).Create(
		context.TODO(),
		cronJob,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// Update Update a cronJob in the client
func (c *CronJob) Update(cronJob *batchv1.CronJob) error {
	if cronJob == nil {
		return errorCronJobEmpty
	}
	_, err := c.client.CronJobs(c.CurrentNamespace).Update(
		context.TODO(),
		cronJob,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete a cronJob in the client
func (c *CronJob) Delete(cronJobName string) error {
	if cronJobName == "" {
		return errorNameEmpty
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := c.client.CronJobs(c.CurrentNamespace).Delete(
		context.TODO(),
		cronJobName,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Get Get a cronJob from the client
func (c *CronJob) Get(cronJobName string) (*batchv1.CronJob, error) {
	if cronJobName == "" {
		return nil, errorNameEmpty
	}
	cronJob, err := c.client.CronJobs(c.CurrentNamespace).Get(
		context.TODO(),
		cronJobName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}

	return cronJob, nil
}

// List List all cronJobs in a namespace
func (c *CronJob) List() (*batchv1.CronJobList, error) {
	cronJobList, err := c.client.CronJobs(c.CurrentNamespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return cronJobList, nil
}
