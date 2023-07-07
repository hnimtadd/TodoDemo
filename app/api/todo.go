package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gosql/model"
	"gosql/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (sv *ConcreteServer) SendTodo(t model.Todo, ch chan int, wg *sync.WaitGroup) {
	id := uuid.New().String()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	t.Id = id
	t.Last_update = time.Now()
	t.Is_deleted = false
	_, err := sv.db.Collection("todos").InsertOne(ctx, t)
	if err == nil {
		ch <- 1
	}
	wg.Done()
}

func (sv *ConcreteServer) AddRandomTodoBatch(ch chan int, wg *sync.WaitGroup, num int) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer func() {
		cancel()
		wg.Done()
	}()
	var wgg sync.WaitGroup
	contents, err := utils.GetRandomContentBatch(ctx, num, sv.config.RandommerKey)
	if err != nil {
		return
	}
	for _, content := range contents[:len(contents)-1] {
		wgg.Add(1)
		go func(content string) {
			defer wgg.Done()
			id := uuid.New().String()
			todo := model.Todo{
				Id:          id,
				Last_update: time.Now(),
				Is_deleted:  false,
				Content:     content,
				Metadata: model.Metadata{
					Size:   1,
					Format: "text",
					Source: "randomAPi",
				},
			}
			if _, err := sv.db.Collection("todos").InsertOne(ctx, todo); err == nil {
				ch <- 1
			}
		}(content)
	}
	wgg.Wait()
	return
}

func (sv *ConcreteServer) FlushTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Access to flush todo func")
		vars := mux.Vars(r)
		num, err := strconv.Atoi(vars["num"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		}
		con, err := strconv.Atoi(vars["con"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		}
		log.Println("checkpoint")
		var wg sync.WaitGroup
		ch := make(chan int, num*con)
		for i := 0; i < num; i++ {
			wg.Add(1)
			go sv.AddRandomTodoBatch(ch, &wg, con)
			wg.Wait()
			log.Printf("Done batch %d\n", i)
		}
		defer func() {
			fmt.Println("Done flush")
			w.Header().Set("Content-Type", "application/json")
			var response struct {
				Msg string `json:"msg"`
			}
			response.Msg = fmt.Sprintf("Flush done, success: %d", len(ch))
			close(ch)
			json.NewEncoder(w).Encode(response)
			fmt.Println("done, response ", response.Msg)
			return
		}()
	}
}

func (sv *ConcreteServer) CountTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		res, err := sv.db.Collection("todos").CountDocuments(ctx, bson.D{})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int64{"num": res})
	}
}
func (sv *ConcreteServer) SearchOneResult() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t model.Todo
		json.NewDecoder(r.Body).Decode(&t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// convert search query to mongo AND search query
		// in mongo AND query must have query field in format like: ""A" "B" "C"", in which A, B, C is search unit
		searchterm := "\"" + strings.Join(strings.Split(t.Content, " "), "\" \"") + "\""
		filter := bson.D{primitive.E{Key: "$text", Value: bson.D{primitive.E{Key: "$search", Value: searchterm}}}}

		cur := sv.db.Collection("todos").FindOne(ctx, filter)
		if cur.Err() != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"msg": cur.Err().Error()})
			return
		}
		var todo model.Todo
		if err := cur.Decode(&todo); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)
		return
	}
}

func (sv *ConcreteServer) SearchWithFullTextSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t model.Todo
		json.NewDecoder(r.Body).Decode(&t)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "$text", Value: bson.D{primitive.E{Key: "$search", Value: t.Content}}}}
		res, err := sv.db.Collection("todos").Find(ctx, filter)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		}
		var todos []model.Todo
		for res.Next(context.TODO()) {
			var result model.Todo
			if err := res.Decode(&result); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			}
			todos = append(todos, result)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}

func (sv *ConcreteServer) SearchWithMysql() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement text search with mongoDB equivalent to
	}
}
func (sv *ConcreteServer) SearchTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t model.Todo
		json.NewDecoder(r.Body).Decode(&t)
		var buf bytes.Buffer
		query := map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"content": map[string]interface{}{
						"query": t.Content,
					},
				},
			},
		}
		if err := json.NewEncoder(&buf).Encode(query); err != nil {
			log.Fatalf("Error encoding query : %s", err)
		}
		res, err := sv.es.Search(
			sv.es.Search.WithContext(context.Background()),
			sv.es.Search.WithIndex("todos"),
			sv.es.Search.WithBody(&buf),
			sv.es.Search.WithPretty(),
		)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		var rs map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&rs); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}
		result := []model.Todo{}
		for _, hit := range rs["hits"].(map[string]interface{})["hits"].([]interface{}) {
			c := hit.(map[string]interface{})["_source"]
			if err := json.NewEncoder(&buf).Encode(c); err != nil {
				log.Fatalf("error encoding query: %s", err)
			}
			t.Id = hit.(map[string]interface{})["_source"].(map[string]interface{})["id"].(string)
			t.Content = hit.(map[string]interface{})["_source"].(map[string]interface{})["content"].(string)
			result = append(result, t)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (sv *ConcreteServer) AddTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t model.Todo
		json.NewDecoder(r.Body).Decode(&t)
		id := uuid.New().String()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		t.Id = id
		t.Last_update = time.Now()
		t.Is_deleted = false
		_, err := sv.db.Collection("todos").InsertOne(ctx, t)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(t)
	}
}

func (sv *ConcreteServer) GetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		cur, err := sv.db.Collection("todos").Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)
		todos := []model.Todo{}
		for cur.Next(ctx) {
			var todo model.Todo
			err := cur.Decode(&todo)
			if err != nil {
				log.Fatal(err)
			}
			todos = append(todos, todo)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}

func (sv *ConcreteServer) GetTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var t model.Todo
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		res := sv.db.Collection("todos").FindOne(ctx, bson.D{primitive.E{Key: "id", Value: id}})
		if res.Err() != nil {
			log.Fatal(res.Err())
		}
		if err := res.Decode(&t); err != nil {
			if err == mongo.ErrNoDocuments {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(t)
	}
}

func (sv *ConcreteServer) UpdateTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t model.Todo
		json.NewDecoder(r.Body).Decode(&t)
		vars := mux.Vars(r)
		id := vars["id"]
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "id", Value: id}}
		update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "last_update", Value: time.Now()}, primitive.E{Key: "content", Value: t.Content}}}}
		res := sv.db.Collection("todos").FindOneAndUpdate(ctx, filter, update)
		if res.Err() != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatal(res.Err())
			return
		}
		if err := res.Decode(&t); err != nil {
			if err == mongo.ErrNoDocuments {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(t)
	}
}

func (sv *ConcreteServer) DeleteTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "id", Value: id}, primitive.E{Key: "is_deleted", Value: false}}
		update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "last_update", Value: time.Now()}, primitive.E{Key: "is_deleted", Value: true}}}}
		var updatedDocument bson.M
		err := sv.db.Collection("todos").FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedDocument)
	}
}
