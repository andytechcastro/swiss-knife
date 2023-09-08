package builders

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// Job struct for build a job
type Job struct {
	Name                    string
	Namespace               string
	Labels                  map[string]string
	Annotations             map[string]string
	MatchLabels             map[string]string
	PodTemplate             corev1.PodTemplateSpec
	BackOffLimit            *int32
	TTLSecondsAfterFinished *int32
	Job                     *batchv1.Job
	JobTemplate             *batchv1.JobTemplateSpec
}

// NewJobBuilder return a job builder
func NewJobBuilder(name string) *Job {
	return &Job{
		Name: name,
	}
}

// SetNamespace set job namespace
func (j *Job) SetNamespace(namespace string) *Job {
	j.Namespace = namespace
	return j
}

// SetLabels set job labels
func (j *Job) SetLabels(labels map[string]string) *Job {
	j.Labels = labels
	return j
}

// SetAnnotations sets job annotations
func (j *Job) SetAnnotations(annotation map[string]string) *Job {
	j.Annotations = annotation
	return j
}

// SetMatchLabels sets job match labels
func (j *Job) SetMatchLabels(matchLabels map[string]string) *Job {
	j.MatchLabels = matchLabels
	return j
}

// SetBackOffLimit set backoff limits in the job
func (j *Job) SetBackOffLimit(backOffLimit int32) *Job {
	j.BackOffLimit = &backOffLimit
	return j
}

// SetTTLSecondsAfterFinished set ttl seconds after finished Job
func (j *Job) SetTTLSecondsAfterFinished(ttlSecondsAfterFinished int32) *Job {
	j.TTLSecondsAfterFinished = &ttlSecondsAfterFinished
	return j
}

// SetPodTemplate set pod template for deployment
func (j *Job) SetPodTemplate(podTemplate corev1.PodTemplateSpec) *Job {
	j.PodTemplate = podTemplate
	return j
}

// Build build a job
func (j *Job) Build() *batchv1.Job {
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:        j.Name,
			Namespace:   j.Namespace,
			Labels:      j.Labels,
			Annotations: j.Annotations,
		},
		Spec: batchv1.JobSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: j.MatchLabels,
			},
			Template:                j.PodTemplate,
			BackoffLimit:            j.BackOffLimit,
			TTLSecondsAfterFinished: j.TTLSecondsAfterFinished,
		},
	}
	j.Job = job
	return job
}

// BuildTemplate build a job
func (j *Job) BuildTemplate() *batchv1.JobTemplateSpec {
	job := &batchv1.JobTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      j.Labels,
			Annotations: j.Annotations,
		},
		Spec: batchv1.JobSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: j.MatchLabels,
			},
			Template:                j.PodTemplate,
			BackoffLimit:            j.BackOffLimit,
			TTLSecondsAfterFinished: j.TTLSecondsAfterFinished,
		},
	}
	j.JobTemplate = job
	return job
}

// ToYaml convert deployment struct to kubernetes yaml
func (j *Job) ToYaml() []byte {
	yaml, err := yaml.Marshal(j.Job)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
