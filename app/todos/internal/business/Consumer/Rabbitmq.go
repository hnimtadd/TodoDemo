package server

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"todosService/config"
	"todosService/internal/business"
	representer "todosService/internal/business/Representer"
	"todosService/protoc"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqConsumer struct {
	Conn     *amqp.Connection
	Queue    string
	Repo     business.TodosRepository
	Config   config.RabbitmqConfig
	Handlers map[string]func(msg amqp.Delivery) (*amqp.Publishing, error)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func NewRabiitMqServer(config config.RabbitmqConfig, repo *business.TodosRepository) business.TodoConsumer {
	sv := &RabbitmqConsumer{
		Config: config,
		Repo:   *repo,
	}
	if err := sv.Setup(); err != nil {
		log.Fatal(err)
	}
	return sv
}

func (sv *RabbitmqConsumer) Setup() error {
	sv.SetupHandlers()
	conn, err := amqp.Dial(sv.Config.Source)
	failOnError(err, "Failed to connect to RabbitMQ")
	sv.Conn = conn
	return nil
}

func (sv *RabbitmqConsumer) Serve(queueName string) error {
	// Declare a channel for consumer to listen
	ch, err := sv.Conn.Channel()
	failOnError(err, "Failed to open a channel")

	// Declare a queue for to listen
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	// Setup QoS, prefect count and size
	err = ch.Qos(
		2,
		0,
		false,
	)
	failOnError(err, "Failed to set QoS")

	// get msgs channel and process message in that channel
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer name
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    //args
	)

	failOnError(err, "Failed to register a consumer")
	var forever chan struct{}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for msg := range msgs {
			action := msg.Type
			handler, ok := sv.Handlers[action]
			if !ok {
				// err := msg.Nack(false, true)
				// failOnError(err, "Failed to nack a message")
				err = msg.Reject(false)
				failOnError(err, "Failed to reject a message")
				log.Println("Don't have any handler")
				continue
			}

			result, err := handler(msg)
			if err != nil {
				// err := msg.Nack(false, false)
				// failOnError(err, "Failed to nack a message")
				err = msg.Reject(false)
				failOnError(err, "Failed to reject a message")
				log.Println("Handler err")
				continue
			}

			err = ch.PublishWithContext(
				ctx,
				"",
				msg.ReplyTo,
				false,
				false,
				*result,
			)
			failOnError(err, "Failed to publish a message")
			err = msg.Ack(false)
			failOnError(err, "Failed to ack a message")
		}
	}()
	log.Printf(" [*] Ready for request")
	<-forever
	return nil
}

func makeMessage(msg amqp.Delivery, v any) (*amqp.Publishing, error) {
	body, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	publish := amqp.Publishing{
		ContentType:   msg.ContentType,
		CorrelationId: msg.CorrelationId,
		Body:          body,
	}
	return &publish, nil
}

func (sv *RabbitmqConsumer) SetupHandlers() {
	handlers := make(map[string]func(amqp.Delivery) (*amqp.Publishing, error))
	handlers["insert"] = sv.AddTodoHandler
	handlers["get"] = sv.GetTodosHandler
	handlers["delete"] = sv.DeleteTodoHandler
	sv.Handlers = handlers
}

func (sv *RabbitmqConsumer) AddTodoHandler(msg amqp.Delivery) (*amqp.Publishing, error) {
	// extract information from message body
	log.Println("received handler")
	// time.Sleep(time.Second * 4)
	var req protoc.AddTodoRequest
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return nil, err
	}

	reporeq, err := representer.AddTodoTransform(&req)
	if err != nil {
		return nil, err
	}
	log.Println("checkpoint")
	// add todo with repo
	reporsp, err := sv.Repo.AddTodo(reporeq)
	if err != nil {
		return nil, err
	}
	log.Println("checkpoint")

	rsp, err := representer.AddTodoDeTransfrom(reporsp)
	if err != nil {
		return nil, err
	}
	log.Println("checkpoint")
	log.Println(rsp)

	body, err := json.Marshal(&rsp)
	if err != nil {
		return nil, err
	}
	log.Println("checkpoint")

	// should reply message
	return &amqp.Publishing{
		ContentType:   msg.ContentType,
		CorrelationId: msg.CorrelationId,
		Body:          body,
	}, nil
	// return &publish, nil
}

func (sv *RabbitmqConsumer) GetTodosHandler(msg amqp.Delivery) (*amqp.Publishing, error) {
	// Get todos from repository
	var req protoc.GetTodosRequest
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("get todos with size= %d, offset = %d", req.Size, req.Offset)
	reporeq, err := representer.GetTodoTransfrom(&req)
	if err != nil {
		return nil, err
	}

	reporsp, err := sv.Repo.GetTodos(reporeq)
	if err != nil {
		return nil, err
	}

	rsp, err := representer.GetTodoDeTransfrom(reporsp)
	if err != nil {
		return nil, err
	}

	// Marshall todos to []byte
	body, err := json.Marshal(&rsp)
	if err != nil {
		return nil, err
	}

	// shoule return publish message
	publish := amqp.Publishing{
		ContentType:   msg.ContentType,
		CorrelationId: msg.CorrelationId,
		Body:          body,
	}
	return &publish, nil
}

func (sv *RabbitmqConsumer) DeleteTodoHandler(msg amqp.Delivery) (*amqp.Publishing, error) {
	var req protoc.DeleteTodoRequest
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("checkpoint")

	repoReq, err := representer.DeleteTodoTransform(&req)
	if err != nil {
		return nil, err
	}

	repoRsp, err := sv.Repo.DeleteTodo(repoReq)
	if err != nil {
		return nil, err
	}

	rsp, err := representer.DeleteTodoDetransform(repoRsp)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(&rsp)
	if err != nil {
		return nil, err
	}

	return &amqp.Publishing{
		ContentType:   msg.ContentType,
		CorrelationId: msg.CorrelationId,
		Body:          body,
	}, nil

}

//	type UpdateBody struct {
//		Id    string `json:"id"`
//		Field string `json:"field"`
//		Value string `json:"value"`
//	}
//
//	func (sv *RabbitmqConsumer) UpdateTodoHandler(msg amqp.Delivery) (*amqp.Publishing, error) {
//		body := msg.Body
//		var req UpdateBody
//		if err := json.Unmarshal(body, &req); err != nil {
//			return nil, err
//		}
//		todo, err := sv.Repo.UpdateTodo(req.Id, req.Field, req.Value)
//		if err != nil {
//			return nil, err
//		}
//
//		body, err = json.Marshal(todo)
//		if err != nil {
//			return nil, err
//		}
//		publish := amqp.Publishing{
//			ContentType:   msg.ContentType,
//			CorrelationId: msg.CorrelationId,
//			Body:          body,
//		}
//		return &publish, nil
//	}
//
// // type GetBody struct {
// // 	Id string `json:"id"`
// // }
// //
// // func (sv *RabbitmqConsumer) GetTodoHandler() any {
// // 	return func(msg rabbitmq.Delivery) rabbitmq.Action {
// // 		body := msg.Body
// // 		var getBodyRequest GetBody
// // 		if err := json.Unmarshal(body, &getBodyRequest); err != nil {
// // 			return rabbitmq.NackRequeue
// // 		}
// // 		todo, err := sv.Repo.GetTodo(getBodyRequest.Id)
// // 		if err != nil {
// // 			return rabbitmq.NackRequeue
// // 		}
// // 		fmt.Println(todo)
// // 		return rabbitmq.Ack
// // 	}
// // }
// //
// // type DeleteBody struct {
// // 	Id string `json:"id"`
// // }
// //
