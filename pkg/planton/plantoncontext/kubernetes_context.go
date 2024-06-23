package plantoncontext

import (
	"github.com/pkg/errors"
	code2cloudenvironmentmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/model"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	plantoncommonsapiresourcemodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	plantonkubeprovidercredential "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/credentials"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/planton/input/apiresource"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/planton/input/credentials"
	plantoncloudpulumisdkkubernetes "github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/automation/provider/kubernetes"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type KubeContext struct {
	state *KubeContextState
	ctx   pulumi.Context
}

type KubeContextState struct {
	Spec   *KubeContextSpec
	Status *KubeContextStatus
}

type KubeContextSpec struct {
	WorkSpaceDir       string
	Labels             map[string]string
	KubeProvider       *kubernetes.Provider
	Metadata           *plantoncommonsapiresourcemodel.ApiResourceMetadata
	EnvironmentInfo    *code2cloudenvironmentmodel.ApiResourceEnvironmentInfo
	IsIngressEnabled   bool
	IngressType        kubernetesworkloadingresstype.KubernetesWorkloadIngressType
	EndpointDomainName string
	EnvDomainName      string
	InternalHostname   string
	ExternalHostname   string
	KubeLocalEndpoint  string
	KubeServiceName    string
}

type KubeContextStatus struct {
	AddedResources *AddedResources
}

type AddedResources struct {
	Namespace                   *kubernetescorev1.Namespace
	LoadBalancerExternalService *kubernetescorev1.Service
	LoadBalancerInternalService *kubernetescorev1.Service
}

func NewKubeContext(ctx *pulumi.Context, stackInput protoreflect.Message) (*KubeContext, error) {

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

	environmentInfo, err := apiresource.ExtractApiResourceEnvironmentInfo(stackInput, &code2cloudenvironmentmodel.ApiResourceEnvironmentInfo{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get api resource environment info")
	}

	var state = &KubeContextSpec{
		KubeProvider:    kubernetesProvider,
		Metadata:        metadata,
		EnvironmentInfo: environmentInfo,
		//IsIngressEnabled   bool
		//IngressType        kubernetesworkloadingresstype.KubernetesWorkloadIngressType
		//EndpointDomainName protostring
		//EnvDomainName      protostring
		//InternalHostname   protostring
		//ExternalHostname   protostring
		//KubeLocalEndpoint  protostring
		//KubeServiceName    protostring
	}
}

func (ctx *KubeContext) WithValue(key, val any) *KubeContext {
	newCtx := &KubeContext{
		ctx:   ctx.ctx,
		state: ctx.state,
		Log:   ctx.Log,
	}
	newCtx.ctx = context.WithValue(newCtx.ctx, key, val)
	return newCtx
}

func (ctx *KubeContext) Value(key any) any {
	return ctx.ctx.Value(key)
}
