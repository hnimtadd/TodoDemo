package api

import (
	"crypto/tls"
	"gosql/config"
	"gosql/server"
	"log"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConcreteServer struct {
	db     *mongo.Database
	router *mux.Router
	es     *elasticsearch.Client
	config config.Config
}

func NewServer(db *mongo.Database, router *mux.Router, es *elasticsearch.Client, config config.Config) server.Server {
	return &ConcreteServer{db: db, router: router, es: es, config: config}
}

func (sv *ConcreteServer) getTLS() (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func (sv *ConcreteServer) Run(port string) error {
	cert, err := sv.getTLS()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	tlsConfig := tls.Config{
		Certificates: []tls.Certificate{*cert},
	}
	sre := &http.Server{
		Handler:      sv.router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		TLSConfig:    &tlsConfig,
	}
	log.Fatal(sre.ListenAndServeTLS("", ""))
	return nil
}

func (sv *ConcreteServer) HelloServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello world"))
	}
}

func (sv *ConcreteServer) MethodNotAlllow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Method not allow"))
	}
}

func (sv *ConcreteServer) NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<h1>Not found this page</h1>
			<h2> welcome to todos page</h2>
			<p> The page you entered doesn't exists in the system</p>
			`))
	}

}

func (sv *ConcreteServer) SetupRoute() error {
	router := sv.router
	router.HandleFunc("/api/todos", sv.GetTodos()).Methods("GET")
	router.HandleFunc("/api/todos/{id}", sv.GetTodo()).Methods("GET")
	router.HandleFunc("/api/todos", sv.AddTodo()).Methods("POST")
	router.HandleFunc("/api/todos/{id}", sv.UpdateTodo()).Methods("PATCH")
	router.HandleFunc("/api/todos/{id}", sv.DeleteTodo()).Methods("DELETE")
	router.HandleFunc("/api/search/todos", sv.SearchTodo()).Methods("GET")
	router.HandleFunc("/api/search/v2/todos", sv.SearchWithMysql()).Methods("GET")
	router.HandleFunc("/api/count/todos", sv.CountTodos()).Methods("GET")
	router.HandleFunc("/api/search/v3/todos", sv.SearchWithFullTextSearch()).Methods("GET")
	router.HandleFunc("/api/search/v4/todo", sv.SearchOneResult()).Methods("GET")
	router.HandleFunc("/api/flush/todos/{num}/{con}", sv.FlushTodo()).Methods("POST")
	router.HandleFunc("/", sv.HelloServer()).Methods("POST", "GET", "PUT")
	router.MethodNotAllowedHandler = sv.MethodNotAlllow()
	router.NotFoundHandler = sv.NotFoundHandler()
	return nil
}
