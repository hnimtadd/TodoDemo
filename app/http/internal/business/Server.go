package business

type Server interface {
	Serve(port string) error
}
