package builders

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// Deployment is the base for work with deployments
type Deployment struct {
	Name        string
	Namespace   string
	Replicas    *int32
	Labels      map[string]string
	Annotations map[string]string
	MatchLabels map[string]string
	PodTemplate corev1.PodTemplateSpec
	Deployment  *appsv1.Deployment
}

// NewDeploymentBuilder return a deployment structr
func NewDeploymentBuilder() *Deployment {
	return &Deployment{}
}

// SetName Set name for deployment
func (d *Deployment) SetName(name string) *Deployment {
	d.Name = name
	return d
}

// SetNamespace Set namespace for deployment
func (d *Deployment) SetNamespace(namespace string) *Deployment {
	d.Namespace = namespace
	return d
}

// SetReplicas Set replicas for deployment
func (d *Deployment) SetReplicas(replicas int32) *Deployment {
	d.Replicas = &replicas
	return d
}

// SetLabels Set labels for deployment
func (d *Deployment) SetLabels(labels map[string]string) *Deployment {
	d.Labels = labels
	return d
}

// SetAnnotations Set Annotations for deployment
func (d *Deployment) SetAnnotations(annotations map[string]string) *Deployment {
	d.Annotations = annotations
	return d
}

// SetMatchLabels Set match labels
func (d *Deployment) SetMatchLabels(matchLabels map[string]string) *Deployment {
	d.MatchLabels = matchLabels
	return d
}

// SetPodTemplate set pod template for deployment
func (d *Deployment) SetPodTemplate(podTemplate corev1.PodTemplateSpec) *Deployment {
	d.PodTemplate = podTemplate
	return d
}

// Build Build de deployment interface
func (d *Deployment) Build() *appsv1.Deployment {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:        d.Name,
			Namespace:   d.Namespace,
			Labels:      d.Labels,
			Annotations: d.Annotations,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: d.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: d.MatchLabels,
			},
			Template: d.PodTemplate,
		},
	}
	d.Deployment = deployment
	return d.Deployment
}

// ToYaml convert deployment struct to kubernetes yaml
func (d *Deployment) ToYaml() []byte {
	yaml, err := yaml.Marshal(d.Deployment)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
