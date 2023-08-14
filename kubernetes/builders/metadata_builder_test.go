package builders_test

import (
	"testing"

	"github.com/andytechcastro/swiss-knife/kubernetes/builders"
	"github.com/stretchr/testify/assert"
)

func initMetadata() *builders.Metadata {
	metadata := builders.NewMetadata("my-service")
	metadata.SetNamespace("default")
	return metadata
}

func TestMetadataToMap(t *testing.T) {
	metadata := initMetadata()
	metadataMap := metadata.ToMap()
	assert.Equal(t, map[string]interface{}{"name": "my-service", "namespace": "default"}, metadataMap)
}

func TestMetadataToMapComplete(t *testing.T) {
	metadata := initMetadata()
	metadata.SetAnnotations(map[string]string{
		"annotation": "my-annotation",
	})
	metadata.SetLabels(map[string]string{
		"label": "my-label",
	})
	metadataMap := metadata.ToMap()
	assert.Equal(
		t,
		map[string]interface{}{
			"name":        "my-service",
			"namespace":   "default",
			"labels":      map[string]interface{}{"label": "my-label"},
			"annotations": map[string]interface{}{"annotation": "my-annotation"},
		},
		metadataMap,
	)
}
