package service

import (
	"fmt"
	"github.com/TwiN/go-color"
	"strconv"
	"strings"
)

type dockerCompose struct {
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	ContainerName string   `yaml:"container_name"`
	Ports         []string `yaml:"ports"`
	DependsOn     []string `yaml:"depends_on"`
	K8sName       string   `yaml:"-"`
	K8sNamespace  string   `yaml:"-"`
	K8sPort       int      `yaml:"-"`
}

func (s Service) FromPort() int {
	var fromPort int

	for _, p := range s.Ports {
		ports := strings.Split(p, ":")
		switch {
		case len(ports) == 1 && strconv.Itoa(s.K8sPort) == ports[0]:
			break
		case len(ports) == 2 && ports[1] == strconv.Itoa(s.K8sPort):
			fromPort, _ = strconv.Atoi(ports[0])
			break
		}
	}

	if fromPort == 0 {
		message := fmt.Sprintf("Failed to match a forwarding port for service %s that maps to the k8s destination %d", s.ContainerName, s.K8sPort)
		fmt.Println(color.Ize(color.Red, message))
		panic(message)
	}

	return fromPort
}

func (s Service) ToPort() int {
	return s.K8sPort
}

func (s Service) Name() string {
	return "service/" + s.K8sName
}

func (s Service) Namespace() string {
	return s.K8sNamespace
}
