package builders_test

import (
	"testing"

	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func initServiceAccount() *corev1.ServiceAccount {
	serviceAccount := corev1.NewServiceAccountBuilder("my-service-account")
	serviceAccount.SetLabels(map[string]string{"label": "my-label"}).
		SetAnnotations(map[string]string{"annotation": "my-annotation"})
	return serviceAccount
}

func TestBuildServiceAccount(t *testing.T) {
	serviceAccount := initServiceAccount()
	buildedServiceAccount := serviceAccount.Build()
	assert.Equal(t, "my-service-account", buildedServiceAccount.GetName())
	assert.Equal(
		t,
		map[string]string{"label": "my-label"},
		buildedServiceAccount.Labels,
	)
	// buildedServiceAccount.Object["metadata"].(map[string]interface{})["labels"],
	assert.Equal(
		t,
		map[string]string{"annotation": "my-annotation"},
		buildedServiceAccount.Annotations,
	)
}

func TestServiceAccountToYaml(t *testing.T) {
	serviceAccount := initServiceAccount()
	serviceAccount.Build()
	yamlServiceAccount := serviceAccount.ToYaml()
	interfaceResult := map[string]interface{}{
		"metadata": map[string]interface{}{
			"annotations": map[string]interface{}{
				"annotation": "my-annotation",
			},
			"creationTimestamp": interface{}(nil),
			"labels": map[string]interface{}{
				"label": "my-label",
			},
			"name": "my-service-account",
		},
	}
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlResult), string(yamlServiceAccount))
}
