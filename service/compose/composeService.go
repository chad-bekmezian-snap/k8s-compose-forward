package compose

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/chad-bekmezian-snap/k8s-compose-forward/file/compose"
	"strconv"
	"strings"
)

type Service struct {
	Spec         compose.ComposeServiceSpec
	K8sName      string
	K8sNamespace string
	K8sPort      int
}

func (s Service) FromPort() int {
	var fromPort int

	for _, p := range s.Spec.Ports {
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
		message := fmt.Sprintf("Failed to match a forwarding port for service %s that maps to the k8s destination %d", s.Spec.ContainerName, s.K8sPort)
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
