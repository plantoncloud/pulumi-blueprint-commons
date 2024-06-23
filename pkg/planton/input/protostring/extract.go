package protostring

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

// ExtractStringValue extracts a string field from a protobuf message given a field path.
func ExtractStringValue(msg protoreflect.Message, fieldPath string) (string, error) {
	fields := strings.Split(fieldPath, ".")
	currentMsg := msg
	var fieldDesc protoreflect.FieldDescriptor

	for _, fieldName := range fields {
		fieldDesc = currentMsg.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		if fieldDesc == nil {
			return "", fmt.Errorf("field '%s' not found in the message", fieldName)
		}

		// If it's the last field in the path, break out to handle it
		if fieldName == fields[len(fields)-1] {
			break
		}

		// The field must be a message type for further traversal
		if fieldDesc.Kind() != protoreflect.MessageKind {
			return "", fmt.Errorf("field '%s' is not a message type", fieldName)
		}

		// Move to the nested message
		currentMsg = currentMsg.Get(fieldDesc).Message()
	}

	// Ensure the final field is of string type
	if fieldDesc.Kind() != protoreflect.StringKind {
		return "", fmt.Errorf("field '%s' is not a string type", fields[len(fields)-1])
	}

	// Get the field value as a protoreflect.Value
	fieldValue := currentMsg.Get(fieldDesc)

	// Return the string value
	return fieldValue.String(), nil
}
