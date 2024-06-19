package loadbalancer

import (
	"github.com/pkg/errors"
	jenkinsservercontextconfig "github.com/plantoncloud/jenkins-server-pulumi-blueprint/pkg/jenkins/contextconfig"
	pulumikubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	v1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) (externalLoadBalancerService *pulumikubernetescorev1.Service,
	internalLoadBalancerService *pulumikubernetescorev1.Service, err error) {
	// Create a Kubernetes Service of type LoadBalancer
	externalLoadBalancerService, err = addExternal(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to add external load balancer")
	}
	internalLoadBalancerService, err = addInternal(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to add internal load balancer")
	}
	return externalLoadBalancerService, internalLoadBalancerService, nil
}

func addExternal(ctx *pulumi.Context) (*pulumikubernetescorev1.Service, error) {
	i := extractInput(ctx)
	addedKubeService, err := pulumikubernetescorev1.NewService(ctx,
		ExternalLoadBalancerServiceName,
		getLoadBalancerServiceArgs(i, ExternalLoadBalancerServiceName, i.ExternalEndpoint), pulumi.Parent(i.Namespace))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create kubernetes service of type load balancer")
	}
	return addedKubeService, nil
}

func addInternal(ctx *pulumi.Context) (*pulumikubernetescorev1.Service, error) {
	i := extractInput(ctx)
	addedKubeService, err := pulumikubernetescorev1.NewService(ctx,
		InternalLoadBalancerServiceName,
		getInternalLoadBalancerServiceArgs(i, i.InternalEndpoint, i.Namespace), pulumi.Parent(i.Namespace))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create kubernetes service of type load balancer")
	}
	return addedKubeService, nil
}

func getInternalLoadBalancerServiceArgs(i *input, hostname string, namespace *pulumikubernetescorev1.Namespace) *pulumikubernetescorev1.ServiceArgs {
	resp := getLoadBalancerServiceArgs(i, InternalLoadBalancerServiceName, hostname)
	resp.Metadata = &v1.ObjectMetaArgs{
		Name:      pulumi.String(InternalLoadBalancerServiceName),
		Namespace: namespace.Metadata.Name(),
		Labels:    namespace.Metadata.Labels(),
		Annotations: pulumi.StringMap{
			"cloud.google.com/load-balancer-type":       pulumi.String("Internal"),
			"planton.cloud/endpoint-domain-name":        pulumi.String(i.EndpointDomainName),
			"external-dns.alpha.kubernetes.io/hostname": pulumi.String(hostname),
		},
	}
	return resp
}

func getLoadBalancerServiceArgs(i *input, serviceName string, hostname string) *pulumikubernetescorev1.ServiceArgs {
	return &pulumikubernetescorev1.ServiceArgs{
		Metadata: &v1.ObjectMetaArgs{
			Name:      pulumi.String(serviceName),
			Namespace: i.Namespace.Metadata.Name(),
			Labels:    i.Namespace.Metadata.Labels(),
			Annotations: pulumi.StringMap{
				"planton.cloud/endpoint-domain-name":        pulumi.String(i.EndpointDomainName),
				"external-dns.alpha.kubernetes.io/hostname": pulumi.String(hostname)}},
		Spec: &pulumikubernetescorev1.ServiceSpecArgs{
			Type: pulumi.String("LoadBalancer"), // Service type is LoadBalancer
			Ports: pulumikubernetescorev1.ServicePortArray{
				&pulumikubernetescorev1.ServicePortArgs{
					Name:       pulumi.String("http"),
					Port:       pulumi.Int(80),
					Protocol:   pulumi.String("TCP"),
					TargetPort: pulumi.String("http"), // This assumes your Jenkins pod has a port named 'http'
				},
			},
			Selector: pulumi.StringMap{
				"app.kubernetes.io/component": pulumi.String("jenkins-controller"),
				"app.kubernetes.io/instance":  i.Namespace.Metadata.Name().Elem(),
			},
		},
	}
}

func addLoadBalancerExternalServiceToContext(existingConfig *jenkinsservercontextconfig.ContextConfig, loadBalancerService *pulumikubernetescorev1.Service) {
	if existingConfig.Status.AddedResources == nil {
		existingConfig.Status.AddedResources = &jenkinsservercontextconfig.AddedResources{
			LoadBalancerExternalService: loadBalancerService,
		}
		return
	}
	existingConfig.Status.AddedResources.LoadBalancerExternalService = loadBalancerService
}

func addLoadBalancerInternalServiceToContext(existingConfig *jenkinsservercontextconfig.ContextConfig, loadBalancerService *pulumikubernetescorev1.Service) {
	if existingConfig.Status.AddedResources == nil {
		existingConfig.Status.AddedResources = &jenkinsservercontextconfig.AddedResources{
			LoadBalancerInternalService: loadBalancerService,
		}
		return
	}
	existingConfig.Status.AddedResources.LoadBalancerInternalService = loadBalancerService
}
