package service

import (
	"fmt"
	"github.com/TwiN/go-color"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

var printLine = fmt.Println

func discardPrintln(_ ...any) (n int, err error) { return 0, nil }

// Load returns a map of docker-compose service names to the details needed to start up their dependencies.
func Load(dockerComposePath string) (map[string]Service, error) {
	if isSilent := os.Getenv("FORWARD_SILENT"); isSilent == "true" {
		printLine = discardPrintln
	}

	var parsedFile *dockerCompose
	var err error
	if parsedFile, err = loadDockerCompose(dockerComposePath); err != nil {
		return nil, err
	}

	k8sServices, err := k8sServices()
	if err != nil {
		return nil, err
	}

	cleanServices(parsedFile.Services, k8sServices)
	printLine()
	printLine()
	printLine()
	printLine()

	return parsedFile.Services, nil
}

func GetBashCompletions(path string) (string, error) {
	compose, err := loadDockerCompose(path)
	if err != nil {
		return "", err
	}

	resultArr := make([]string, 0, len(compose.Services))
	for name, _ := range compose.Services {
		resultArr = append(resultArr, name)
	}

	return strings.Join(resultArr, " "), nil
}

func loadDockerCompose(path string) (*dockerCompose, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(file)
	var parsedFile dockerCompose
	if err := decoder.Decode(&parsedFile); err != nil {
		return nil, err
	}

	return &parsedFile, nil
}

func cleanServices(services map[string]Service, k8sServices k8sSvcs) {
	for serviceName, svc := range services {
		k8Svc, match := k8sServices.FindServiceByClosestMatchingName(serviceName)
		if !match {
			printLine(color.Ize(color.Blue, fmt.Sprintf("Unable to find a k8s service matching the name %s. Skipping.", serviceName)))
			delete(services, serviceName)
			continue
		}

		printLine(color.Ize(color.Green, fmt.Sprintf("Matching %s -> k8s/%s", serviceName, k8Svc.Detail.Name)))
		svc.K8sName = k8Svc.Detail.Name
		svc.K8sNamespace = k8Svc.Detail.Namespace
		svc.K8sPort = k8Svc.Spec.Ports[0].Port
		services[serviceName] = svc
	}
}
