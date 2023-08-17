package repository

import (
	"context"
	"fmt"
	"log"
	"time"
	"todosService/config"
	"todosService/internal/business"
	"todosService/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoRepository struct {
	db     *mongo.Database
	config *config.MongoConfig
}

func NewMongRepository(config config.MongoConfig) (business.TodosRepository, error) {
	instance := &MongoRepository{
		config: &config,
	}
	if err := instance.initSetup(); err != nil {
		return nil, err
	}
	return instance, nil
}

func (repo *MongoRepository) initSetup() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI(repo.config.Source).SetAuth(
		options.Credential{
			AuthSource: repo.config.AuthSource,
			Username:   repo.config.Username,
			Password:   repo.config.Password,
		},
	).SetTLSConfig(nil)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	repo.db = client.Database(repo.config.Database)
	return nil
}

func (repo *MongoRepository) AddTodo(req *business.AddTodoRequest) (*business.AddTodoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	_, err := repo.db.Collection("todos").InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp := business.AddTodoResponse{
		Todo: req.Todo,
	}
	return &rsp, nil
}

func (repo *MongoRepository) GetTodos(req *business.GetTodosRequest) (*business.GetTodoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	log.Printf("Enter get todos, with size = %d, offset = %d", req.Size, req.Offset)
	// query1 := bson.A{bson.D{{"$limit", size}}, bson.D{{"$skip", offset}}}
	// facetStage := bson.D{{"$facet", bson.D{{"stage", query1}}}}
	// cur, err := repo.db.Collection("todos").Aggregate(ctx, mongo.Pipeline{facetStage})

	cur, err := repo.db.Collection("todos").Find(ctx, bson.D{}, options.Find().SetLimit(int64(req.Size)).SetSkip(int64(req.Offset)))
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	todos := []model.Todo{}
	for cur.Next(ctx) {
		var todo model.Todo
		if err := cur.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	log.Printf("get todos, len: %d", len(todos))
	rsp := business.GetTodoResponse{
		Result: todos,
	}
	return &rsp, nil
}
func (repo *MongoRepository) UpdateTodo(id string, field, value string) (*model.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "last_update", Value: time.Now()}, primitive.E{Key: field, Value: value}}}}
	res := repo.db.Collection("todos").FindOneAndUpdate(ctx, filter, update)
	if res.Err() != nil {
		return nil, res.Err()
	}
	var todo model.Todo
	if err := res.Decode(&todo); err != nil {
		return nil, err
	}
	return &todo, nil
}
func (repo *MongoRepository) GetTodo(id string) (*model.Todo, error) {
	var todo model.Todo
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := repo.db.Collection("todos").FindOne(ctx, bson.D{primitive.E{Key: "id", Value: id}})
	if res.Err() != nil {
		return nil, res.Err()
	}
	if err := res.Decode(&todo); err != nil {
		return nil, err
	}
	return &todo, nil
}
func (repo *MongoRepository) DeleteTodo(req *business.DeleteTodoRequest) (*business.DeleteTodoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "id", Value: req.Id}, primitive.E{Key: "is_deleted", Value: false}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "last_update", Value: time.Now()}, primitive.E{Key: "is_deleted", Value: true}}}}
	res := repo.db.Collection("todos").FindOneAndUpdate(ctx, filter, update)
	if err := res.Err(); err != nil {
		return nil, err
	}
	return &business.DeleteTodoResponse{Result: fmt.Sprintf("Delete todo with id: %s success", req.Id)}, nil
}
