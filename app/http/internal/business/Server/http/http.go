package http

import (
	"httpService/config"
	"httpService/internal/business"
	publisher "httpService/internal/business/Server/http/Publisher"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FiberSever struct {
	app       *fiber.App
	publisher publisher.Publisher
}

func NewFiberServer(publisher publisher.Publisher, config config.ServerConfig) business.Server {
	server := &FiberSever{
		publisher: publisher,
	}
	if err := server.Setup(); err != nil {
		log.Fatalf("error: %v", err)
	}
	return server
}

func (sv *FiberSever) Serve(port string) error {
	for _, route := range sv.app.GetRoutes() {
		log.Printf("%s - %s", route.Method, route.Path)
	}
	if err := sv.app.Listen(":" + port); err != nil {
		log.Fatalf("error: %v", err)
	}
	return nil
}

func (sv *FiberSever) Setup() error {
	config := fiber.Config{
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}
	sv.app = fiber.New(config)
	sv.InitRoute()
	return nil
}

func (sv *FiberSever) InitRoute() {
	sv.app.Post("/api/v1/todos", sv.AddTodoHandler())
	sv.app.Get("/api/v1/todos", sv.GetTodosHandler())
	sv.app.Delete("/api/v1/todos/:id", sv.DeleteTodoHandler())
}
