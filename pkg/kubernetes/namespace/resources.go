package namespace

import (
	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Resources(ctx *pulumi.Context, stackInput protoreflect.Message) (*kubernetescorev1.Namespace, error) {
	i, err := extractInput(ctx, stackInput)
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract input for namespace resource")
	}
	namespace, err := addNamespace(ctx, i)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add namespace")
	}
	return namespace, nil
}

func addNamespace(ctx *pulumi.Context, i *input) (*kubernetescorev1.Namespace, error) {
	ns, err := kubernetescorev1.NewNamespace(ctx, i.Metadata.Id, &kubernetescorev1.NamespaceArgs{
		ApiVersion: pulumi.String("v1"),
		Kind:       pulumi.String("Namespace"),
		Metadata: metav1.ObjectMetaPtrInput(&metav1.ObjectMetaArgs{
			Name:   pulumi.String(i.Metadata.Id),
			Labels: pulumi.ToStringMap(i.Metadata.Labels),
		}),
	}, pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "5s", Update: "5s", Delete: "5s"}),
		pulumi.Provider(i.KubernetesProvider))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to add %s namespace", i.Metadata.Id)
	}
	return ns, nil
}
