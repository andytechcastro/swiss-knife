package actions

import (
	batchInterface "k8s.io/client-go/kubernetes/typed/batch/v1"
)

// BatchV1 struct for access to batchv1 api
type BatchV1 struct {
	Job     *Job
	CronJob *CronJob
}

// NewBatchV1 return a batch v1 api
func NewBatchV1(client batchInterface.BatchV1Interface) *BatchV1 {
	return &BatchV1{
		Job:     NewJobAction(client),
		CronJob: NewCronJobAction(client),
	}
}
