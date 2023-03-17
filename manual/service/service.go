package service

type Service struct {
	K8sName       string
	K8sNamespace  string
	PortFlagName  string
	DefaultPort   int
	ForwardToPort int
}

func (s Service) FromPort() int {
	return s.DefaultPort
}

func (s Service) ToPort() int {
	return s.ForwardToPort
}

func (s Service) Name() string {
	return s.K8sName
}

func (s Service) Namespace() string {
	return s.K8sNamespace
}

type Configurable struct {
	DefaultPort int
	Service     Service
}
