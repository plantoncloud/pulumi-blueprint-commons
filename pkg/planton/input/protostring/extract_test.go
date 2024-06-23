package protostring

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	plantonapiresourcetesting "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/testing/message"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestExtractStringValue(t *testing.T) {
	// Helper to create protoreflect.Message
	createTestMessage := func(name string) protoreflect.Message {
		msg := &plantonapiresourcetesting.ApiResourceKubernetesTest{
			Metadata: &model.ApiResourceMetadata{
				Name: name,
			},
		}
		return msg.ProtoReflect()
	}

	// 1. Valid Extraction
	t.Run("ValidExtraction", func(t *testing.T) {
		expectedValue := "test-postgres"
		msg := createTestMessage(expectedValue)

		result, err := ExtractStringValue(msg, "metadata.name")
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, result)
	})

	//2. Field Not Found
	t.Run("FieldNotFound", func(t *testing.T) {
		expectedValue := ""
		msg := createTestMessage(expectedValue)

		result, err := ExtractStringValue(msg, "metadata.name2")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'name2' not found")
		assert.Equal(t, "", result)
	})

	// 3. Invalid Field Type
	t.Run("InvalidFieldType", func(t *testing.T) {
		msg := createTestMessage("")

		result, err := ExtractStringValue(msg, "name")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'name' not found")
		assert.Equal(t, "", result)
	})

	// 5. Empty Field Path
	t.Run("EmptyFieldPath", func(t *testing.T) {
		msg := createTestMessage("empty_path")

		result, err := ExtractStringValue(msg, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field '' not found")
		assert.Equal(t, "", result)
	})
}
