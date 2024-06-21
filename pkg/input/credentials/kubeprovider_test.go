package credentials

import (
	postgreskubernetesstatemodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/postgreskubernetes/model"
	postgreskubernetesstackmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/postgreskubernetes/stack/kubernetes/model"
	apiresourcecommonsmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	iacv1sjmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/credentials"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"
)

func TestExtractAndConvertCredentials(t *testing.T) {
	// Define a helper function to create a protoreflect.Message
	createTestMessage := func(cred *iacv1sjmodel.KubernetesProviderCredential) protoreflect.Message {
		msg := &postgreskubernetesstackmodel.PostgresKubernetesStackInput{
			CredentialsInput: cred,
		}
		return msg.ProtoReflect()
	}

	// 1. Valid Extraction
	t.Run("ValidExtraction", func(t *testing.T) {
		expectedCred := &iacv1sjmodel.KubernetesProviderCredential{
			KubeconfigBase64: "valid_credential",
		}
		msg := createTestMessage(expectedCred)
		extractedCred := &iacv1sjmodel.KubernetesProviderCredential{}

		result, err := ExtractKubeProvider(msg, extractedCred)
		assert.NoError(t, err)
		assert.Equal(t, expectedCred.KubeconfigBase64, result.KubeconfigBase64)
	})

	// 2. Field Not Found
	t.Run("FieldNotFound", func(t *testing.T) {
		msg := &postgreskubernetesstatemodel.PostgresKubernetes{
			Metadata: &apiresourcecommonsmodel.ApiResourceMetadata{
				Name:    "",
				Id:      "",
				Labels:  nil,
				Version: nil,
			},
		}
		extractedCred := &iacv1sjmodel.KubernetesProviderCredential{}

		result, err := ExtractKubeProvider(msg.ProtoReflect(), extractedCred)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'credentials_input' not found in the message")
		assert.Empty(t, result.KubeconfigBase64)
	})

	// 3. Invalid Field Type
	t.Run("InvalidFieldType", func(t *testing.T) {
		msg := &postgreskubernetesstackmodel.PostgresKubernetesStackInput{
			// Simulate invalid field type (e.g., an int instead of a message)
			ResourceInput: &postgreskubernetesstatemodel.PostgresKubernetes{},
		}
		extractedCred := &iacv1sjmodel.KubernetesProviderCredential{}

		result, err := ExtractKubeProvider(msg.ProtoReflect(), extractedCred)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "'credentials_input' is invalid or nil")
		assert.Empty(t, result.KubeconfigBase64)
	})

	// 4. Invalid Field Value
	t.Run("InvalidFieldValue", func(t *testing.T) {
		msg := &postgreskubernetesstackmodel.PostgresKubernetesStackInput{
			CredentialsInput: nil, // Simulate nil or invalid field value
		}
		extractedCred := &iacv1sjmodel.KubernetesProviderCredential{}

		result, err := ExtractKubeProvider(msg.ProtoReflect(), extractedCred)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "'credentials_input' is invalid or nil")
		assert.Empty(t, result.KubeconfigBase64)
	})

	// 5. Marshal/Unmarshal Errors
	t.Run("MarshalUnmarshalErrors", func(t *testing.T) {
		// Create a dynamic message with incompatible types to simulate marshal/unmarshal error
		invalidMsg := &anypb.Any{
			TypeUrl: "invalid/type",
			Value:   []byte("invalid_data"),
		}
		extractedCred := &iacv1sjmodel.KubernetesProviderCredential{}

		result, err := ExtractKubeProvider(invalidMsg.ProtoReflect(), extractedCred)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'credentials_input' not found in the message")
		assert.Empty(t, result.KubeconfigBase64)
	})
}
