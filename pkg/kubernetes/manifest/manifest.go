package manifest

import (
	"github.com/pkg/errors"
	pulumik8syaml "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/printers"
	"os"
	"path/filepath"
)

func CreateAndDeploy(ctx *pulumi.Context, k8sManifestsLoc string, k8sResource runtime.Object, resourceName string, opts ...pulumi.ResourceOption) error {
	if err := os.MkdirAll(filepath.Dir(k8sManifestsLoc), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", filepath.Dir(k8sManifestsLoc))
	}
	newFile, err := os.Create(k8sManifestsLoc)
	if err != nil {
		return errors.Wrapf(err, "failed to create deployment manifest %s", k8sManifestsLoc)
	}
	y := printers.YAMLPrinter{}
	defer newFile.Close()
	if err := y.PrintObj(k8sResource, newFile); err != nil {
		return errors.Wrapf(err, "failed to write deployment manifest %s", k8sManifestsLoc)
	}
	log.Debugf("successfully created %s", k8sManifestsLoc)
	_, err = pulumik8syaml.NewConfigFile(ctx, resourceName, &pulumik8syaml.ConfigFileArgs{File: k8sManifestsLoc}, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to add gateway manifest")
	}
	return nil
}
