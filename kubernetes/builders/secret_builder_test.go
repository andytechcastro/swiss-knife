package builders_test

import (
	"encoding/base64"
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/yaml"
)

func initSecret() *builders.Secret {
	secret := builders.NewSecretBuilder("my-secret")
	secret.SetNamespace("default")
	return secret
}

func TestBuildSecret(t *testing.T) {
	secret := initSecret()
	secret.SetStringData(map[string]string{"secret": "my-secret"})
	buildedSecret := secret.Build()
	assert.Equal(t, "my-secret", buildedSecret.Name)
	assert.Equal(t, "my-secret", buildedSecret.StringData["secret"])
}

func TestBuildSecretBase64(t *testing.T) {
	secret := initSecret()
	pass := "my-secret"
	passEnc := base64.StdEncoding.EncodeToString([]byte(pass))
	secret.SetData(map[string][]byte{
		"secret": []byte(passEnc),
	},
	).SetLabels(map[string]string{
		"label": "my-label",
	}).SetAnnotations(map[string]string{
		"annotation": "my-annotation",
	})
	buildedSecret := secret.Build()
	passDec, _ := base64.StdEncoding.DecodeString(passEnc)
	assert.Equal(t, "my-secret", buildedSecret.Name)
	assert.Equal(t, "my-secret", string(passDec))
	assert.Equal(t, "my-label", buildedSecret.Labels["label"])
	assert.Equal(t, "my-annotation", buildedSecret.Annotations["annotation"])
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
