package loadbalancer

import (
	"github.com/pkg/errors"
	pulumikubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	pulumikubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	ExternalServiceName             = "ingress-external-lb"
	InternalServiceName             = "ingress-internal-lb"
	ExternalExternalNameServiceName = "ingress-external-external-dns"
	InternalExternalNameServiceName = "ingress-internal-external-dns"
)

func Resources(ctx *pulumi.Context, internalMetaData *pulumikubernetesmetav1.ObjectMetaArgs,
	externalMetaData *pulumikubernetesmetav1.ObjectMetaArgs,
	serviceSpecArgs *pulumikubernetescorev1.ServiceSpecArgs, opts ...pulumi.ResourceOption) (
	externalLoadBalancerService *pulumikubernetescorev1.Service, internalLoadBalancerService *pulumikubernetescorev1.Service, err error) {
	externalLoadBalancerService, err = pulumikubernetescorev1.NewService(ctx, ExternalServiceName, getLoadBalancerServiceArgs(externalMetaData, serviceSpecArgs), opts...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create kubernetes service for external load balancer")
	}

	internalLoadBalancerService, err = pulumikubernetescorev1.NewService(ctx, InternalServiceName, getLoadBalancerServiceArgs(internalMetaData, serviceSpecArgs), opts...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create kubernetes service for external load balancer")
	}
	return externalLoadBalancerService, internalLoadBalancerService, nil
}

func getLoadBalancerServiceArgs(metaData *pulumikubernetesmetav1.ObjectMetaArgs, serviceSpecArgs *pulumikubernetescorev1.ServiceSpecArgs) *pulumikubernetescorev1.ServiceArgs {
	return &pulumikubernetescorev1.ServiceArgs{
		Metadata: metaData,
		Spec:     serviceSpecArgs,
	}
}
