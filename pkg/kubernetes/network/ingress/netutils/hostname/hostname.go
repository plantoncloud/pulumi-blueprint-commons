package hostname

import (
	"fmt"
	"github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/network/dns"
)

func GetIngressInternalHostname(mongodbClusterId, environmentName, endpointDomainName string) string {
	return fmt.Sprintf("%s.%s-internal.%s", mongodbClusterId, environmentName, endpointDomainName)
}

func GetIngressExternalHostname(resourceId, environmentName, endpointDomainName string) string {
	return fmt.Sprintf("%s.%s.%s", resourceId, environmentName, endpointDomainName)
}

func GetKubeEndpoint(resourceName, namespace string) string {
	return fmt.Sprintf("%s.%s.%s", GetKubeServiceName(resourceName), namespace, dns.DefaultDomain)
}

func GetKubeServiceName(resourceName string) string {
	return fmt.Sprintf(resourceName)
}
