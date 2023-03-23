package file

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ServiceName string

type Compose struct {
	// Services is a map of a service name or key to its compose specification.
	Services map[string]ComposeServiceSpec `yaml:"services"`
}

type ComposeServiceSpec struct {
	ContainerName string   `yaml:"container_name"`
	Ports         []string `yaml:"ports"`
	DependsOn     []string `yaml:"depends_on"`
}

func ParseCompose(filePath string) (Compose, error) {
	var result Compose
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}
