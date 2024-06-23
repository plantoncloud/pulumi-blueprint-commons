package plantoncontext

import (
	code2cloudenvironmentmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/model"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	plantoncommonsapiresourcemodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Context struct {
	state *ContextState
	ctx   pulumi.Context
}

type ContextState struct {
	Spec   *Spec
	Status *Status
}

type Spec struct {
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

type Status struct {
	AddedResources *AddedResources
}

type AddedResources struct {
	Namespace                   *kubernetescorev1.Namespace
	LoadBalancerExternalService *kubernetescorev1.Service
	LoadBalancerInternalService *kubernetescorev1.Service
}

func NewContext(ctx *pulumi.Context) (*Context, error) {

}

func (ctx *Context) WithValue(key, val any) *Context {
	newCtx := &Context{
		ctx:   ctx.ctx,
		state: ctx.state,
		Log:   ctx.Log,
	}
	newCtx.ctx = context.WithValue(newCtx.ctx, key, val)
	return newCtx
}

func (ctx *Context) Value(key any) any {
	return ctx.ctx.Value(key)
}
