package service

type Service struct {
	Name          string
	Namespace     string
	PortFlagName  string
	DefaultPort   uint
	ForwardToPort uint
}

type Configurable struct {
	DefaultPort uint
	Service     Service
}
