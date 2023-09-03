package builders

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// CronJob struct for build a cronJob
type CronJob struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
	Schedule    string
	JobTemplate batchv1.JobTemplateSpec
	CronJob     *batchv1.CronJob
}

// NewCronJobBuilder return a cronJob builder
func NewCronJobBuilder(name string) *CronJob {
	return &CronJob{
		Name: name,
	}
}

// SetNamespace set cronJob namespace
func (c *CronJob) SetNamespace(namespace string) *CronJob {
	c.Namespace = namespace
	return c
}

// SetLabels set cronJob labels
func (c *CronJob) SetLabels(labels map[string]string) *CronJob {
	c.Labels = labels
	return c
}

// SetAnnotations sets cronJob annotations
func (c *CronJob) SetAnnotations(annotation map[string]string) *CronJob {
	c.Annotations = annotation
	return c
}

// SetSchedule set job schedule
func (c *CronJob) SetSchedule(schedule string) *CronJob {
	c.Schedule = schedule
	return c
}

// SetJobTemplate set pod template for deployment
func (c *CronJob) SetJobTemplate(cronJobTemplate batchv1.JobTemplateSpec) *CronJob {
	c.JobTemplate = cronJobTemplate
	return c
}

// Build build a cronJob
func (c *CronJob) Build() *batchv1.CronJob {
	cronJob := &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:        c.Name,
			Namespace:   c.Namespace,
			Labels:      c.Labels,
			Annotations: c.Annotations,
		},
		Spec: batchv1.CronJobSpec{
			Schedule:    c.Schedule,
			JobTemplate: c.JobTemplate,
		},
	}
	c.CronJob = cronJob
	return cronJob
}

// ToYaml convert deployment struct to kubernetes yaml
func (c *CronJob) ToYaml() []byte {
	yaml, err := yaml.Marshal(c.CronJob)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
