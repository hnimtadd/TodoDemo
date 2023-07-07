package main

import (
	"context"
	"fmt"
	"gosql/api"
	"gosql/config"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	esconfig := elasticsearch.Config{
		Addresses:    []string{config.ElasticURL},
		DisableRetry: true,
	}
	es, err := elasticsearch.NewClient(esconfig)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(config.Dsn)

	clientOpts := options.Client().ApplyURI(config.Dsn).SetAuth(
		options.Credential{
			AuthSource: config.DbAuthsource,
			Username:   config.DbUsername,
			Password:   config.DbPassword,
		},
	).SetTLSConfig(nil)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Initialized connect to mysql")
	router := mux.NewRouter()
	db := client.Database(config.Database)
	sv := api.NewServer(db, router, es, *config)
	sv.SetupRoute()
	sv.Run(config.ServerPort)
}
