package protomessage

import (
	plantonstatemodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/postgreskubernetes/model"
	plantonstackmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/postgreskubernetes/stack/kubernetes/model"
	plantonkubecredentials "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/credentials"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestExtractProtoMessage(t *testing.T) {
	// Helper to create protoreflect.Message
	createTestMessage := func(cred *plantonkubecredentials.KubernetesProviderCredential) protoreflect.Message {
		msg := &plantonstackmodel.PostgresKubernetesStackInput{
			CredentialsInput: cred,
		}
		return msg.ProtoReflect()
	}

	// 1. Valid Extraction
	t.Run("ValidExtraction", func(t *testing.T) {
		expectedCred := &plantonkubecredentials.KubernetesProviderCredential{KubeconfigBase64: "valid_credential"}
		msg := createTestMessage(expectedCred)
		extractedCred := &plantonkubecredentials.KubernetesProviderCredential{}

		result, err := ExtractProtoMessage(msg, "credentials_input", extractedCred)
		assert.NoError(t, err)
		assert.Equal(t, expectedCred.KubeconfigBase64, result.KubeconfigBase64)
	})

	// 2. Field Not Found
	t.Run("FieldNotFound", func(t *testing.T) {
		msg := &plantonkubecredentials.KubernetesProviderCredential{}
		extractedCred := &plantonkubecredentials.KubernetesProviderCredential{}

		result, err := ExtractProtoMessage(msg.ProtoReflect(), "NonExistentField", extractedCred)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'NonExistentField' not found")
		assert.Empty(t, result.KubeconfigBase64)
	})

	// 3. Invalid Field Type
	t.Run("InvalidFieldType", func(t *testing.T) {
		msg := &plantonstackmodel.PostgresKubernetesStackInput{
			// Simulate invalid field type, assuming ResourceInput is not a message type
			CredentialsInput: &plantonkubecredentials.KubernetesProviderCredential{},
			ResourceInput:    &plantonstatemodel.PostgresKubernetes{},
		}
		extractedCred := &plantonkubecredentials.KubernetesProviderCredential{}

		result, err := ExtractProtoMessage(msg.ProtoReflect(), "resource_input", extractedCred)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "type mismatch: expected credentials.KubernetesProviderCredential, got *model.PostgresKubernetes")
		assert.Empty(t, result.KubeconfigBase64)
	})

	// 4. Invalid Field Value
	t.Run("InvalidFieldValue", func(t *testing.T) {
		msg := &plantonstackmodel.PostgresKubernetesStackInput{
			CredentialsInput: nil, // Simulate nil field value
		}
		extractedCred := &plantonkubecredentials.KubernetesProviderCredential{}

		result, err := ExtractProtoMessage(msg.ProtoReflect(), "credentials_input", extractedCred)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'credentials_input' is invalid or nil")
		assert.Empty(t, result.KubeconfigBase64)
	})

	// 5. Nested Fields
	t.Run("NestedFields", func(t *testing.T) {
		msg := &plantonstackmodel.PostgresKubernetesStackInput{
			CredentialsInput: &plantonkubecredentials.KubernetesProviderCredential{KubeconfigBase64: "nested_credential"},
		}
		extractedCred := &plantonkubecredentials.KubernetesProviderCredential{}

		result, err := ExtractProtoMessage(msg.ProtoReflect(), "credentials_input.kubeconfig_base64", extractedCred)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field 'kubeconfig_base64' is not a message type")
		assert.Empty(t, result.KubeconfigBase64)
	})

	// 6. Empty Field Path
	t.Run("EmptyFieldPath", func(t *testing.T) {
		msg := createTestMessage(&plantonkubecredentials.KubernetesProviderCredential{KubeconfigBase64: "empty_path"})
		extractedCred := &plantonkubecredentials.KubernetesProviderCredential{}

		result, err := ExtractProtoMessage(msg, "", extractedCred)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "field '' not found")
		assert.Empty(t, result.KubeconfigBase64)
	})
}
