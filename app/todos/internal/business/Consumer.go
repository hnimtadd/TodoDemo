package business

type TodoConsumer interface {
	Setup() error
	Serve(string) error
}
