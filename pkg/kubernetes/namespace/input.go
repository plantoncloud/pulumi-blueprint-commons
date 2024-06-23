package namespace

import (
	"github.com/pkg/errors"
	plantoncommonsapiresourcemodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	plantonkubeprovidercredential "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/credentials"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/planton/input/apiresource"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/planton/input/credentials"
	plantoncloudpulumisdkkubernetes "github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/automation/provider/kubernetes"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type input struct {
	KubernetesProvider *pulumikubernetes.Provider
	Metadata           *plantoncommonsapiresourcemodel.ApiResourceMetadata
}

func extractInput(ctx *pulumi.Context, stackInput protoreflect.Message) (*input, error) {
	kubeProviderCredential, err := credentials.ExtractKubeProvider(stackInput, &plantonkubeprovidercredential.KubernetesProviderCredential{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract kubernetes provider credentials from stack input")
	}
	kubernetesProvider, err := plantoncloudpulumisdkkubernetes.GetWithStackCredentials(ctx, kubeProviderCredential)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get kubernetes provider with stack kubernetes provider credentials")
	}
	metadata, err := apiresource.ExtractMetadata(stackInput, &plantoncommonsapiresourcemodel.ApiResourceMetadata{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get api resource metadata")
	}

	return &input{
		KubernetesProvider: kubernetesProvider,
		Metadata:           metadata,
	}, nil
}
