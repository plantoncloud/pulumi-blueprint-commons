package protostring

import (
	plantonkubernetesstatemodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/model"
	plantonapiresourcetesting "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/testing/message"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestExtractStringValue(t *testing.T) {
	// Helper to create protoreflect.Message
	createTestMessage := func(isIngressEnabled bool) protoreflect.Message {
		msg := &plantonapiresourcetesting.ApiResourceKubernetesTest{
			Spec: &plantonapiresourcetesting.ApiResourceKubernetesSpec{
				Ingress: &plantonkubernetesstatemodel.ApiResourceIngressSpec{
					IsEnabled: isIngressEnabled,
				},
			},
		}
		return msg.ProtoReflect()
	}

	// 1. Valid Extraction
	t.Run("ValidExtraction", func(t *testing.T) {
		expectedValue := true
		msg := createTestMessage(expectedValue)

		result, err := ExtractBoolValue(msg, "spec.ingress.is_enabled")
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, result)
	})

	//2. Field Not Found
	t.Run("FieldNotFound", func(t *testing.T) {
		expectedValue := false
		msg := createTestMessage(expectedValue)

		result, err := ExtractBoolValue(msg, "spec.ingress.is_enabled_not_found")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'is_enabled_not_found' not found")
		assert.Equal(t, false, result)
	})

	// 3. Invalid Field Type
	t.Run("InvalidFieldType", func(t *testing.T) {
		msg := createTestMessage(false)

		result, err := ExtractBoolValue(msg, "invalid_field_type")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'invalid_field_type' not found")
		assert.Equal(t, false, result)
	})

	// 5. Empty Field Path
	t.Run("EmptyFieldPath", func(t *testing.T) {
		msg := createTestMessage(false)

		result, err := ExtractBoolValue(msg, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field '' not found")
		assert.Equal(t, false, result)
	})
}
