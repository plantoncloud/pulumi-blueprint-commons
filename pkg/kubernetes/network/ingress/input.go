package ingress

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/network"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	IngressType kubernetesworkloadingresstype.KubernetesWorkloadIngressType
}

func extractInput(ctx *pulumi.Context) *input {
	var ingressInput = ctx.Value(network.IngressInputKey).(network.IngressInput)

	return &input{
		IngressType: ingressInput.IngressType,
	}
}
