package protostring

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

// ExtractBoolValue extracts a boolean field from a protobuf message given a field path.
func ExtractBoolValue(msg protoreflect.Message, fieldPath string) (bool, error) {
	fields := strings.Split(fieldPath, ".")
	currentMsg := msg
	var fieldDesc protoreflect.FieldDescriptor

	for _, fieldName := range fields {
		fieldDesc = currentMsg.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		if fieldDesc == nil {
			return false, fmt.Errorf("field '%s' not found in the message", fieldName)
		}

		// If it's the last field in the path, break out to handle it
		if fieldName == fields[len(fields)-1] {
			break
		}

		// The field must be a message type for further traversal
		if fieldDesc.Kind() != protoreflect.MessageKind {
			return false, fmt.Errorf("field '%s' is not a message type", fieldName)
		}

		// Move to the nested message
		currentMsg = currentMsg.Get(fieldDesc).Message()
	}

	// Ensure the final field is of boolean type
	if fieldDesc.Kind() != protoreflect.BoolKind {
		return false, fmt.Errorf("field '%s' is not a boolean type", fields[len(fields)-1])
	}

	// Get the field value as a protoreflect.Value
	fieldValue := currentMsg.Get(fieldDesc)

	// Return the boolean value
	return fieldValue.Bool(), nil
}
