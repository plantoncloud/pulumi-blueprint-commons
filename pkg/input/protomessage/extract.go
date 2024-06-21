package protomessage

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
	"strings"
)

func ExtractProtoMessage[T proto.Message](msg protoreflect.Message, fieldPath string, target T) (T, error) {
	// Split the field path into components
	fields := strings.Split(fieldPath, ".")

	// Traverse the field path to find the target field
	currentMsg := msg
	var fieldDesc protoreflect.FieldDescriptor

	for _, fieldName := range fields {
		fieldDesc = currentMsg.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		if fieldDesc == nil {
			return target, fmt.Errorf("field '%s' not found in the message", fieldName)
		}

		// If it's the last field in the path, break out to handle it
		if fieldName == fields[len(fields)-1] {
			break
		}

		// If not the last field, the field must be a message type
		if fieldDesc.Kind() != protoreflect.MessageKind {
			return target, fmt.Errorf("field '%s' is not a message type", fieldName)
		}

		// Move to the nested message
		currentMsg = currentMsg.Get(fieldDesc).Message()
	}

	// Ensure the final field is a message type
	if fieldDesc.Kind() != protoreflect.MessageKind {
		return target, fmt.Errorf("field '%s' is not a message type", fields[len(fields)-1])
	}

	// Get the field value as a protoreflect.Value
	fieldValue := currentMsg.Get(fieldDesc)

	// Check if the field value is valid
	if !fieldValue.Message().IsValid() {
		return target, fmt.Errorf("field '%s' is invalid or nil", fields[len(fields)-1])
	}

	// Convert protoreflect.Message to proto.Message and then to the target type
	anyMessage := fieldValue.Message().Interface()

	// Type check: Compare the type of the field value with the target type
	if reflect.TypeOf(anyMessage) != reflect.TypeOf(&target).Elem() {
		return target, fmt.Errorf("type mismatch: expected %s, got %s", reflect.TypeOf(target).Elem(), reflect.TypeOf(anyMessage))
	}

	// Use proto.Marshal to convert to bytes and then unmarshal to the target type
	data, err := proto.Marshal(anyMessage)
	if err != nil {
		return target, fmt.Errorf("failed to marshal: %v", err)
	}

	if err := proto.Unmarshal(data, target); err != nil {
		return target, fmt.Errorf("failed to unmarshal to target type: %v", err)
	}

	return target, nil
}
