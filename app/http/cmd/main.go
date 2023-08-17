package main

import (
	"httpService/config"
	"httpService/internal/business/Server/http"
	publisher "httpService/internal/business/Server/http/Publisher"
	"log"
)

func main() {
	serverconfig := config.NewServerConfig(".")
	log.Println(serverconfig)
	rabbitmqconfig := config.NewRabbitMQConfig(".")
	publisher := publisher.NewRabbitMQPublisher(rabbitmqconfig)
	server := http.NewFiberServer(publisher, serverconfig)
	if err := server.Serve(serverconfig.ServerPort); err != nil {
		log.Fatalf("error: %v", err)
	}
}
