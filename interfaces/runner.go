package interfaces

type Runner interface {
	Run() error
	Cancel() error
}

type NamedMessage struct {
	Name   string
	Output []byte
}
