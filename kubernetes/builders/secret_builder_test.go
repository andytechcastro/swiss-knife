package builders_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	//"gopkg.in/yaml.v3"
)

func initSecret() *builders.Secret {
	secret := builders.NewSecretBuilder()
	secret.SetName("my-secret").
		SetNamespace("default")
	return secret
}

func TestBuildSecret(t *testing.T) {
	secret := initSecret()
	secret.SetStringData(map[string]string{"secret": "my-secret"})
	buildedSecret := secret.Build()
	assert.Equal(t, "my-secret", buildedSecret.Name)
	assert.Equal(t, "my-secret", buildedSecret.StringData["secret"])
}

func TestSecretToYaml(t *testing.T) {
	secret := initSecret()
	secret.SetStringData(map[string]string{
		"username": "admin",
		"password": "P@ssw0rd",
	})
	secret.SetType("kubernetes.io/basic-auth").Build()
	yamlSecret := secret.ToYaml()
	interfaceResult := map[string]interface{}(
		map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"creationTimestamp": interface{}(nil),
				"name":              "my-secret",
				"namespace":         "default",
			},
			"type": "kubernetes.io/basic-auth",
			"stringData": map[string]string{
				"username": "admin",
				"password": "P@ssw0rd",
			},
			"immutable": false,
		},
	)
	yamlResult, _ := yaml.Marshal(interfaceResult)
	assert.YAMLEq(t, string(yamlSecret), string(yamlResult))
}
