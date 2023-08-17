package main

import (
	"log"
	"todosService/config"
	server "todosService/internal/business/Consumer"
	repository "todosService/internal/business/Repository"
)

func main() {
	repoconfig := config.NewMongoConfig(".")
	log.Println(repoconfig)
	log.Println("Setting up repository...")
	repo, err := repository.NewMongRepository(repoconfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Println("Setting up client...")
	rabbitConfig := config.NewRabbitmqConfig(".")
	log.Println(rabbitConfig)
	client := server.NewRabiitMqServer(rabbitConfig, &repo)
	log.Println("Serving...")
	if err := client.Serve("todos"); err != nil {
		log.Fatalf("error: %v", err)
	}
}
