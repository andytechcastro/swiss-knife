package actions

import "errors"

var (
	errorDeploymentEmpty     = errors.New("The deployment is empty. Exec a build or a Get")
	errorPodEmpty            = errors.New("The pod is empty. Exec a build or a Get")
	errorServiceEmpty        = errors.New("The service is empty. Exec a build or a Get")
	errorServiceAccountEmpty = errors.New("The service account is empty. Exec a build or a Get")
	errorNamespaceEmpty      = errors.New("The namespace is empty. Exec a build or a Get")
	errorNameEmpty           = errors.New("The name is empty")
)
