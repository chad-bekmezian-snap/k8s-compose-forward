package forward

type Service interface {
	FromPort() int
	ToPort() int
	Name() string
	Namespace() string
}
