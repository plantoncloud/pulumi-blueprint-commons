package virtualservice

import (
	"fmt"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/manifest"
	"istio.io/client-go/pkg/apis/networking/v1beta1"

	"path/filepath"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, workspaceDir string, virtualService *v1beta1.VirtualService, opts ...pulumi.ResourceOption) error {
	resourceName := fmt.Sprintf("virtual-service-%s", virtualService.Name)
	manifestPath := filepath.Join(workspaceDir, fmt.Sprintf("%s.yaml", resourceName))

	if err := manifest.CreateAndDeploy(ctx, manifestPath, virtualService, resourceName, opts...); err != nil {
		return errors.Wrapf(err, "failed to create or deploy %s manifest file", manifestPath)
	}
	return nil
}
