package compose

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/chad-bekmezian-snap/k8s-compose-forward/file/compose"
	"github.com/chad-bekmezian-snap/k8s-compose-forward/service/k8s"
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

	var parsedFile compose.Compose
	var err error
	if parsedFile, err = compose.ParseCompose(dockerComposePath); err != nil {
		return nil, err
	}

	k8sServices, err := k8s.ListServices()
	if err != nil {
		return nil, err
	}

	result := mapComposeToK8s(parsedFile, k8sServices)
	printLine()
	printLine()
	printLine()
	printLine()

	return result, nil
}

func GetBashCompletions(path string) (string, error) {
	compose, err := compose.ParseCompose(path)
	if err != nil {
		return "", err
	}

	resultArr := make([]string, 0, len(compose.Services))
	for name, _ := range compose.Services {
		resultArr = append(resultArr, name)
	}

	return strings.Join(resultArr, " "), nil
}

func mapComposeToK8s(composeFile compose.Compose, k8sServices k8s.ServiceSlice) map[string]Service {
	result := make(map[string]Service, len(composeFile.Services))

	for serviceName, svc := range composeFile.Services {
		k8Svc, match := k8sServices.FindServiceByClosestMatchingName(serviceName)
		if !match {
			printLine(color.Ize(color.Blue, fmt.Sprintf("Unable to find a k8s service matching the name %s. Skipping.", serviceName)))
			continue
		}

		printLine(color.Ize(color.Green, fmt.Sprintf("Matching %s -> k8s/%s", serviceName, k8Svc.Detail.Name)))
		result[serviceName] = Service{
			Spec:         svc,
			K8sName:      k8Svc.Detail.Name,
			K8sNamespace: k8Svc.Detail.Namespace,
			K8sPort:      k8Svc.Spec.Ports[0].Port,
		}
	}

	return result
}
