package loadbalancer

import (
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/network"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	ExternalLoadBalancerServiceName             = "ingress-external-lb"
	InternalLoadBalancerServiceName             = "ingress-internal-lb"
	ExternalLoadBalancerExternalNameServiceName = "ingress-external-external-dns"
	InternalLoadBalancerExternalNameServiceName = "ingress-internal-external-dns"
)

type input struct {
	ResourceId         string
	ResourceName       string
	Namespace          *kubernetescorev1.Namespace
	ExternalEndpoint   string
	InternalEndpoint   string
	EndpointDomainName string
	ServiceName        string
}

func extractInput(ctx *pulumi.Context) *input {
	var ingressInput = ctx.Value(network.IngressInputKey).(network.IngressInput)

	return &input{
		ResourceId:         ingressInput.ResourceId,
		ResourceName:       ingressInput.ResourceName,
		Namespace:          ingressInput.Namespace,
		EndpointDomainName: ingressInput.EndpointDomainName,
	}
}
