package kubernetes_test

import (
	"fmt"
	kube "kodikas-libraries-go/kubernetes"
	"testing"
)

func TestDeploymentToYaml(t *testing.T) {
	container := kube.NewContainerBuilder()
	container.SetName("test")
	container.SetImage("nginx")
	container.SetTag("1")
	container.SetPort(80)
	deployment := kube.NewDeploymentBuilder()
	deployment.AddContainer(container.Build())
	deployment.SetName("test")
	deployment.SetNamespace("test")
	deployment.SetReplicas(3)
	deployment.SetLabels(map[string]string{"test": "testing"})
	deployment.SetAnnotations(map[string]string{"annotation": "testAnnotation"})
	deployment.SetMatchLabels(map[string]string{"test": "testingmatch"})
	deployment.Build()
	yaml := deployment.ToYaml()
	fmt.Println(string(yaml))
}

func TestServiceToYaml(t *testing.T) {
	service := kube.NewServiceBuilder()
	service.SetName("my-service")
	service.Build()
	yaml := service.ToYaml()
	fmt.Println(string(yaml))
}
