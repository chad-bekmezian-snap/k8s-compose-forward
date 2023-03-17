package service

type dockerCompose struct {
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	ContainerName string   `yaml:"container_name"`
	Ports         []string `yaml:"ports"`
	DependsOn     []string `yaml:"depends_on"`
	K8sName       string
	K8sNamespace  string
	K8sPort       int
}
