package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	actionsCoreV1 "github.com/andytechcastro/swiss-knife/kubernetes/actions/core/v1"
	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

func initSecret() *actionsCoreV1.Secret {
	info := map[string]string{
		"secret1": "firstPass",
		"secret2": "my-secret",
		"secret3": "ultra-secret",
		"secret4": "weakPass",
	}
	objects := []runtime.Object{}
	for name, secretValue := range info {
		secret := corev1.NewSecretBuilder(name)
		buildedSecret := secret.SetNamespace("default").
			SetType("kubernetes.io/basic-auth").
			SetStringData(map[string]string{
				"username": name,
				"password": secretValue,
			}).Build()
		objects = append(objects, buildedSecret)
	}
	objectsDynamic := []runtime.Object{}
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	client := fake.NewSimpleClientset(objects...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).CoreV1.Secret
	actions.Namespace("default")
	return actions
}

func TestGetSecret(t *testing.T) {
	actions := initSecret()
	secret, _ := actions.Get("secret3")
	assert.Equal(t, "secret3", secret.Name)
	assert.Equal(t, "secret3", secret.StringData["username"])
	assert.Equal(t, "ultra-secret", secret.StringData["password"])
}

func TestCreateSecret(t *testing.T) {
	actions := initSecret()
	secret := corev1.NewSecretBuilder("secret5")
	buildedSecret := secret.SetNamespace("default").
		SetType("kubernetes.io/basic-auth").
		SetStringData(map[string]string{
			"username": "secret5",
			"password": "strongPass",
		}).Build()
	actions.Create(buildedSecret)
	newSecret, _ := actions.Get("secret5")
	secrets, _ := actions.List()
	assert.Equal(t, "secret5", newSecret.Name)
	assert.Equal(t, "secret5", secret.StringData["username"])
	assert.Equal(t, "strongPass", secret.StringData["password"])
	assert.Equal(t, 5, len(secrets.Items))
}

func TestUpdateSecret(t *testing.T) {
	actions := initSecret()
	secret, _ := actions.Get("secret3")
	immutable := true
	secret.Immutable = &immutable
	actions.Update(secret)
	updatedSecret, _ := actions.Get("secret3")
	assert.Equal(t, "secret3", updatedSecret.Name)
	assert.Equal(t, "secret3", updatedSecret.StringData["username"])
	assert.Equal(t, "ultra-secret", updatedSecret.StringData["password"])
	assert.Equal(t, immutable, *updatedSecret.Immutable)
}

func TestDeleteSecret(t *testing.T) {
	actions := initSecret()
	actions.Delete("secret4")
	secrets, _ := actions.List()
	assert.Equal(t, 3, len(secrets.Items))
	for _, secret := range secrets.Items {
		assert.NotEqual(t, "secret4", secret.Name)
	}
}

func TestListSecret(t *testing.T) {
	actions := initSecret()
	secrets, _ := actions.List()
	assert.Equal(t, 4, len(secrets.Items))
}
