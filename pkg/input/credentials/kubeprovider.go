package credentials

import (
	"fmt"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/input/protomessage"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ExtractKubeProvider[T proto.Message](msg protoreflect.Message, target T) (T, error) {
	credentials, err := protomessage.ExtractProtoMessage(msg, "credentials_input", target)
	if err != nil {
		return target, fmt.Errorf("failed to extract kube provider credentials: %v", err)
	}
	return credentials, nil
}
