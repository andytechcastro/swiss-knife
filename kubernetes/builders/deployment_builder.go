package builders

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
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
	Containers  []apiv1.Container
	Deployment  *appsv1.Deployment
}

// NewDeploymentBuilder return a deployment structr
func NewDeploymentBuilder() *Deployment {
	return &Deployment{}
}

// SetName Set name for deployment
func (d *Deployment) SetName(name string) {
	d.Name = name
}

// SetNamespace Set namespace for deployment
func (d *Deployment) SetNamespace(namespace string) {
	d.Namespace = namespace
}

// AddContainer Add new container to deployment
func (d *Deployment) AddContainer(container apiv1.Container) {
	d.Containers = append(d.Containers, container)
}

// SetReplicas Set replicas for deployment
func (d *Deployment) SetReplicas(replicas int32) {
	d.Replicas = &replicas
}

// SetLabels Set labels for deployment
func (d *Deployment) SetLabels(labels map[string]string) {
	d.Labels = labels
}

// SetAnnotations Set Annotations for deployment
func (d *Deployment) SetAnnotations(annotations map[string]string) {
	d.Annotations = annotations
}

// SetMatchLabels Set match labels
func (d *Deployment) SetMatchLabels(matchLabels map[string]string) {
	d.MatchLabels = matchLabels
}

// Build Build de deployment interface
func (d *Deployment) Build() *appsv1.Deployment {
	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
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
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: d.MatchLabels,
				},
				Spec: apiv1.PodSpec{
					Containers: d.Containers,
				},
			},
		},
	}
	d.Deployment = deployment
	return deployment
}

// ToYaml convert deployment struct to kubernetes yaml
func (d *Deployment) ToYaml() []byte {
	yaml, err := yaml.Marshal(d.Deployment)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
