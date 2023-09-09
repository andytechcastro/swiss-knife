package actions

import (
	"context"

	"github.com/andytechcastro/swiss-knife/errors"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreInterface "k8s.io/client-go/kubernetes/typed/core/v1"
)

// Secret struct for action secret
type Secret struct {
	client           coreInterface.CoreV1Interface
	CurrentNamespace string
}

// NewSecretAction get an actions secret
func NewSecretAction(client coreInterface.CoreV1Interface) *Secret {
	return &Secret{
		client:           client,
		CurrentNamespace: "",
	}
}

// Namespace set namespace
func (s *Secret) Namespace(namespace string) *Secret {
	s.CurrentNamespace = namespace
	return s
}

// Create Create a secrets in the client
func (s *Secret) Create(secret *apiv1.Secret) error {
	if secret == nil {
		return errors.GetEmptyError("Secret")
	}
	_, err := s.client.Secrets(s.CurrentNamespace).Create(
		context.TODO(),
		secret,
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// Update Update a secrets in the client
func (s *Secret) Update(secret *apiv1.Secret) error {
	if secret == nil {
		return errors.GetEmptyError("Secret")
	}
	_, err := s.client.Secrets(s.CurrentNamespace).Update(
		context.TODO(),
		secret,
		metav1.UpdateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete a secret in the client
func (s *Secret) Delete(secretName string) error {
	if secretName == "" {
		return errors.GetEmptyError("Name")
	}
	deletePolicy := metav1.DeletePropagationForeground
	err := s.client.Secrets(s.CurrentNamespace).Delete(
		context.TODO(),
		secretName,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Get Get a secret from the client
func (s *Secret) Get(secretName string) (*apiv1.Secret, error) {
	if secretName == "" {
		return nil, errors.GetEmptyError("Name")
	}
	secret, err := s.client.Secrets(s.CurrentNamespace).Get(
		context.TODO(),
		secretName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

// List List all secrets in a namespace
func (s *Secret) List() (*apiv1.SecretList, error) {
	secretList, err := s.client.Secrets(s.CurrentNamespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return secretList, nil
}
