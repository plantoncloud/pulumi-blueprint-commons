package ingress

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	istiocommons "github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/network/ingress/istio"
	loadbalancercommons "github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/network/ingress/loadbalancer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) (newCtx *pulumi.Context, err error) {
	i := extractInput(ctx)
	switch i.IngressType {
	case kubernetesworkloadingresstype.KubernetesWorkloadIngressType_load_balancer:
		ctx, err = loadbalancercommons.Resources(ctx)
		if err != nil {
			return ctx, errors.Wrap(err, "failed to add load balancer resources")
		}
	case kubernetesworkloadingresstype.KubernetesWorkloadIngressType_ingress_controller:
		if err = istiocommons.Resources(ctx); err != nil {
			return ctx, errors.Wrap(err, "failed to add istio resources")
		}
	}
	return ctx, nil
}
