package gateway

import (
	"fmt"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/manifest"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, workspaceDir string, gateway *v1beta1.Gateway, opts ...pulumi.ResourceOption) error {
	resourceName := fmt.Sprintf("gateway-%s", gateway.Name)
	manifestPath := filepath.Join(workspaceDir, fmt.Sprintf("%s.yaml", resourceName))

	if err := manifest.CreateAndDeploy(ctx, manifestPath, gateway, resourceName, opts...); err != nil {
		return errors.Wrapf(err, "failed to create or deploy %s manifest file", manifestPath)
	}
	return nil
}
