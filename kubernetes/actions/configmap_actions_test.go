package actions_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/actions"
	corev1 "github.com/andytechcastro/swiss-knife/kubernetes/builders/core/v1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

func initConfigMaps() *actions.ConfigMap {
	info := map[string]string{
		"configmap1": "default",
		"configmap2": "default",
		"configmap3": "default",
		"configmap4": "default",
	}
	objects := []runtime.Object{}
	for name, namespace := range info {
		cm := corev1.NewConfigMapBuilder(name)
		cms := cm.SetNamespace(namespace).
			SetData(map[string]string{
				"myKey": name,
			}).Build()
		objects = append(objects, cms)
	}
	client := fake.NewSimpleClientset(objects...)
	objectsDynamic := []runtime.Object{}

	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtime.NewScheme(), objectsDynamic...)
	actions := actions.GetActionFilled(client, dynamicClient, &rest.Config{}).ConfigMap
	actions.Namespace("default")
	return actions
}

func TestGetConfigMap(t *testing.T) {
	actions := initConfigMaps()
	configMap, _ := actions.Get("configmap2")
	assert.Equal(t, "configmap2", configMap.Name)
	assert.Equal(t, "default", configMap.Namespace)
	assert.Equal(t, "configmap2", configMap.Data["myKey"])
}

func TestCreateConfigMap(t *testing.T) {
	actions := initConfigMaps()
	configMap := corev1.NewConfigMapBuilder("configmap5")
	buildedConfigMap := configMap.SetNamespace("default").
		SetData(map[string]string{
			"myKey": "newConfigMap",
		}).
		Build()

	actions.Create(buildedConfigMap)
	newConfigMap, _ := actions.Get("configmap5")
	configMaps, _ := actions.List()
	assert.Equal(t, "configmap5", newConfigMap.Name)
	assert.Equal(t, "newConfigMap", newConfigMap.Data["myKey"])
	assert.Equal(t, 5, len(configMaps.Items))
}

func TestUpdateConfigMap(t *testing.T) {
	actions := initConfigMaps()
	configMap, _ := actions.Get("configmap3")
	configMap.Data["newKey"] = "newConfigMap"
	actions.Update(configMap)
	configMapUpdated, _ := actions.Get("configmap3")
	assert.Equal(t, "newConfigMap", configMapUpdated.Data["newKey"])
}

func TestDeleteConfigMap(t *testing.T) {
	actions := initConfigMaps()
	actions.Delete("configmap4")
	services, _ := actions.List()
	assert.Equal(t, 3, len(services.Items))
	for _, service := range services.Items {
		assert.NotEqual(t, "configmap4", service.Name)
	}
}

func TestListConfigMap(t *testing.T) {
	actions := initConfigMaps()
	configMaps, _ := actions.List()
	assert.Equal(t, 4, len(configMaps.Items))
}
