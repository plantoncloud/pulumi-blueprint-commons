package apiresource

import (
	"fmt"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/planton/input/protomessage"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ExtractApiResourceEnvironmentInfo[T proto.Message](msg protoreflect.Message, target T) (T, error) {
	environmentInfo, err := protomessage.ExtractProtoMessage(msg, "resource_input.spec.environment_info", target)
	if err != nil {
		return target, fmt.Errorf("failed to extract environment info: %v", err)
	}
	return environmentInfo, nil
}
