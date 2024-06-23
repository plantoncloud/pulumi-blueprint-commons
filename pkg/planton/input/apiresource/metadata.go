package apiresource

import (
	"fmt"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/planton/input/protomessage"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ExtractMetadata[T proto.Message](msg protoreflect.Message, target T) (T, error) {
	metadata, err := protomessage.ExtractProtoMessage(msg, "resource_input.metadata", target)
	if err != nil {
		return target, fmt.Errorf("failed to extract metadata: %v", err)
	}
	return metadata, nil
}
