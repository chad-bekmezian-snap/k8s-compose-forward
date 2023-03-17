package service

import (
	"fmt"
	"github.com/TwiN/go-color"
	"gopkg.in/yaml.v3"
	"os"
)

// Load returns a map of docker-compose service names to the details needed to start up their dependencies.
func Load(dockerComposePath string) (map[string]Service, error) {
	file, err := os.Open(dockerComposePath)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(file)
	var parsedFile dockerCompose
	if err := decoder.Decode(&parsedFile); err != nil {
		return nil, err
	}

	k8sServices, err := k8sServices()
	if err != nil {
		return nil, err
	}

	cleanServices(parsedFile.Services, k8sServices)
	println()
	println()
	println()
	println()

	return parsedFile.Services, nil
}

func cleanServices(services map[string]Service, k8sServices k8sSvcs) {
	for serviceName, svc := range services {
		k8Svc, match := k8sServices.FindServiceByClosestMatchingName(serviceName)
		if !match {
			fmt.Println(color.Ize(color.Blue, fmt.Sprintf("Unable to find a service in the cloud matching the name %s. Skipping.", serviceName)))
			delete(services, serviceName)
			continue
		}

		fmt.Println(color.Ize(color.Green, fmt.Sprintf("Matching %s -> %s", serviceName, k8Svc.Detail.Name)))
		svc.K8sName = k8Svc.Detail.Name
		svc.K8sNamespace = k8Svc.Detail.Namespace
		svc.K8sPort = k8Svc.Spec.Ports[0].Port
		services[serviceName] = svc
	}
}
