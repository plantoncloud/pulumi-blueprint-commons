package namespace

import (
	"github.com/pkg/errors"
	plantoncommonsapiresourcemodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	plantonkubeprovidercredential "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/credentials"
	apiresource "github.com/plantoncloud/pulumi-blueprint-commons/pkg/input/apiresource"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/input/credentials"
	plantoncloudpulumisdkkubernetes "github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/automation/provider/kubernetes"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Resources(ctx *pulumi.Context, stackInput protoreflect.Message) (*kubernetescorev1.Namespace, error) {
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

	namespace, err := addNamespace(ctx, kubernetesProvider, metadata)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add namespace")
	}
	return namespace, nil
}

func addNamespace(ctx *pulumi.Context, kubernetesProvider *pulumikubernetes.Provider, metadata *plantoncommonsapiresourcemodel.ApiResourceMetadata) (*kubernetescorev1.Namespace, error) {
	ns, err := kubernetescorev1.NewNamespace(ctx, metadata.Id, &kubernetescorev1.NamespaceArgs{
		ApiVersion: pulumi.String("v1"),
		Kind:       pulumi.String("Namespace"),
		Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
			Name:   pulumi.String(metadata.Id),
			Labels: pulumi.ToStringMap(metadata.Labels),
		}),
	}, pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "5s", Update: "5s", Delete: "5s"}),
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to add %s namespace", "i.NamespaceName")
	}
	return ns, nil
}
