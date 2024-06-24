package protoenum

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	plantonkubeclustermodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/model"
	plantontestingapiresourcemessage "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/testing/message"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestExtractEnumValue(t *testing.T) {
	// Helper to create protoreflect.Message
	createTestMessage := func(ingressType kubernetesworkloadingresstype.KubernetesWorkloadIngressType) protoreflect.Message {
		msg := &plantontestingapiresourcemessage.ApiResourceKubernetesTest{
			Spec: &plantontestingapiresourcemessage.ApiResourceKubernetesTestSpec{
				Ingress: &plantonkubeclustermodel.IngressSpec{
					IsEnabled:   true,
					IngressType: ingressType,
				},
			},
		}
		return msg.ProtoReflect()
	}

	// 1. Valid Extraction
	t.Run("ValidExtraction", func(t *testing.T) {
		expectedValue := kubernetesworkloadingresstype.KubernetesWorkloadIngressType_load_balancer
		msg := createTestMessage(expectedValue)

		var target kubernetesworkloadingresstype.KubernetesWorkloadIngressType
		result, err := ExtractEnumValue(msg, "spec.ingress.ingress_type", target)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, result)
	})

	//// 2. Field Not Found
	//t.Run("FieldNotFound", func(t *testing.T) {
	//	msg := dynamicpb.NewMessageType(&descriptorpb.DescriptorProto{}).New().ProtoReflect()
	//
	//	var target ExampleEnum
	//	result, err := ExtractEnumValue(msg, "NonExistentField", &target)
	//	assert.Error(t, err)
	//	assert.Contains(t, err.Error(), "field 'NonExistentField' not found")
	//	assert.Equal(t, ExampleEnum(0), result)
	//})
	//
	//// 3. Invalid Field Type
	//t.Run("InvalidFieldType", func(t *testing.T) {
	//	msg := createTestMessage(ExampleEnum_UNKNOWN)
	//
	//	var target ExampleEnum
	//	result, err := ExtractEnumValue(msg, "InvalidField", &target)
	//	assert.Error(t, err)
	//	assert.Contains(t, err.Error(), "field 'InvalidField' not found")
	//	assert.Equal(t, ExampleEnum(0), result)
	//})
	//
	//// 4. Nested Fields
	//t.Run("NestedFields", func(t *testing.T) {
	//	nestedMsg := &ExampleMessage{Status: ExampleEnum_INACTIVE}
	//	msg := &ExampleMessage{
	//		Status: ExampleEnum_INACTIVE,
	//	}.ProtoReflect()
	//
	//	var target ExampleEnum
	//	result, err := ExtractEnumValue(msg, "Status", &target)
	//	assert.NoError(t, err)
	//	assert.Equal(t, ExampleEnum_INACTIVE, result)
	//})
	//
	//// 5. Empty Field Path
	//t.Run("EmptyFieldPath", func(t *testing.T) {
	//	msg := createTestMessage(ExampleEnum_ACTIVE)
	//
	//	var target ExampleEnum
	//	result, err := ExtractEnumValue(msg, "", &target)
	//	assert.Error(t, err)
	//	assert.Contains(t, err.Error(), "field '' not found")
	//	assert.Equal(t, ExampleEnum(0), result)
	//})
}
