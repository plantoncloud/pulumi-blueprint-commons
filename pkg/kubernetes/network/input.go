package network

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
)

const (
	IngressInputKey = "ingress-input"
)

type IngressInput struct {
	WorkspaceDir       string
	ResourceId         string
	ResourceName       string
	Namespace          *kubernetescorev1.Namespace
	NamespaceName      string
	KubeProvider       *pulumikubernetes.Provider
	Labels             map[string]string
	IsIngressEnabled   bool
	IngressType        kubernetesworkloadingresstype.KubernetesWorkloadIngressType
	EndpointDomainName string
	EnvDomainName      string
}
