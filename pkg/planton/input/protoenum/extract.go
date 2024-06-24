package protoenum

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
	"strings"
)

// ExtractEnumValue extracts an enum field from a protobuf message given a field path and target type.
func ExtractEnumValue[T protoreflect.Enum](msg protoreflect.Message, fieldPath string, target T) (T, error) {
	fields := strings.Split(fieldPath, ".")
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

		// The field must be a message type for further traversal
		if fieldDesc.Kind() != protoreflect.MessageKind {
			return target, fmt.Errorf("field '%s' is not a message type", fieldName)
		}

		// Move to the nested message
		currentMsg = currentMsg.Get(fieldDesc).Message()
	}

	// Ensure the final field is of enum type
	if fieldDesc.Kind() != protoreflect.EnumKind {
		return target, fmt.Errorf("field '%s' is not an enum type", fields[len(fields)-1])
	}

	// Get the field value as a protoreflect.Value
	fieldValue := currentMsg.Get(fieldDesc)

	// Convert the enum value to the target enum type
	enumValue := fieldValue.Enum()
	targetType := reflect.TypeOf(target)
	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}
	enumValueReflected := reflect.ValueOf(enumValue).Convert(targetType)

	return enumValueReflected.Interface().(T), nil
}
