package network

import (
	"github.com/pkg/errors"
	commoningress "github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/network/ingress"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, i *IngressInput) (newCtx *pulumi.Context, err error) {
	ctx = ctx.WithValue(IngressInputKey, *i)
	if !i.IsIngressEnabled || i.EndpointDomainName == "" {
		return ctx, nil
	}
	if ctx, err = commoningress.Resources(ctx); err != nil {
		return ctx, errors.Wrap(err, "failed to add gateway resources")
	}
	return ctx, nil
}
