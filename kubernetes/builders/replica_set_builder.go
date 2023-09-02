package builders

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// ReplicaSet struct for build replicaset
type ReplicaSet struct {
	Name        string
	Namespace   string
	Replicas    *int32
	Labels      map[string]string
	Annotations map[string]string
	PodTemplate corev1.PodTemplateSpec
	ReplicaSet  *appsv1.ReplicaSet
	MatchLabels map[string]string
}

// NewReplicaSetBuilder return a ReplicaSet builder
func NewReplicaSetBuilder() *ReplicaSet {
	return &ReplicaSet{}
}

// SetName set the replicaset name
func (rs *ReplicaSet) SetName(name string) *ReplicaSet {
	rs.Name = name
	return rs
}

// SetNamespace set the namespace of the replicaset
func (rs *ReplicaSet) SetNamespace(namespace string) *ReplicaSet {
	rs.Namespace = namespace
	return rs
}

// SetReplicas set the replicaset replicas
func (rs *ReplicaSet) SetReplicas(replicas int32) *ReplicaSet {
	rs.Replicas = &replicas
	return rs
}

// SetLabels set the replicaset labels
func (rs *ReplicaSet) SetLabels(labels map[string]string) *ReplicaSet {
	rs.Labels = labels
	return rs
}

// SetAnnotations set the replicaset annotations
func (rs *ReplicaSet) SetAnnotations(annotations map[string]string) *ReplicaSet {
	rs.Annotations = annotations
	return rs
}

// SetMatchLabels set the labels selector for the replicaset
func (rs *ReplicaSet) SetMatchLabels(labels map[string]string) *ReplicaSet {
	rs.MatchLabels = labels
	return rs
}

// SetPodTemplate set the pod template for the replicaset
func (rs *ReplicaSet) SetPodTemplate(template corev1.PodTemplateSpec) *ReplicaSet {
	rs.PodTemplate = template
	return rs
}

// Build build the Replicaset
func (rs *ReplicaSet) Build() *appsv1.ReplicaSet {
	replicaSet := &appsv1.ReplicaSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:        rs.Name,
			Namespace:   rs.Namespace,
			Labels:      rs.Labels,
			Annotations: rs.Annotations,
		},
		Spec: appsv1.ReplicaSetSpec{
			Replicas: rs.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: rs.MatchLabels,
			},
			Template: rs.PodTemplate,
		},
	}
	rs.ReplicaSet = replicaSet
	return replicaSet
}

// ToYaml convert deployment struct to kubernetes yaml
func (rs *ReplicaSet) ToYaml() []byte {
	yaml, err := yaml.Marshal(rs.ReplicaSet)
	if err != nil {
		fmt.Println(err)
	}
	return yaml
}
