package publisher

import (
	"context"
	"errors"
	"httpService/config"
	"httpService/utils"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	conn   *amqp.Connection
	config *config.RabbitMQConfig
}

func NewRabbitMQPublisher(config config.RabbitMQConfig) Publisher {
	publisher := &RabbitMQPublisher{
		config: &config,
	}
	if err := publisher.Setup(); err != nil {
		log.Fatalf("error: %v", err)
	}
	return publisher
}

func (pb *RabbitMQPublisher) Setup() error {
	conn, err := amqp.Dial(pb.config.Source)
	if err != nil {
		return err
	}
	pb.conn = conn
	return nil
}

func (pb *RabbitMQPublisher) Publish(req PublishRequest) (*PublishResponse, error) {
	ch, err := pb.conn.Channel()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if err := ch.Confirm(false); err != nil {
		log.Fatalf("error: %v", err)
	}
	q, err := ch.QueueDeclare( // reply queue
		"",
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Replyqueue
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	corrId := utils.RandomString(32)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// TODO: continue implement publish method
	err = ch.PublishWithContext(
		ctx,
		"",
		"todos",
		false,
		false,
		amqp.Publishing{
			Type:          req.Action,
			ContentType:   req.ContentType,
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          req.Body,
		})

	var wg sync.WaitGroup

	result := make(chan PublishResponse, 1)

	// listen to response
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			log.Println("hllo")
			msg, ok := <-msgs
			log.Println("[Msg]", msg)
			if !ok {
				return
			}
			if corrId == msg.CorrelationId {
				rsp := PublishResponse{
					Code: 200,
					Body: []byte("success"),
				}
				log.Println("msg: ", msg)
				result <- rsp
				return
			}
			log.Println("Timeout")
		}
	}()

	chann := make(chan amqp.Confirmation, 1)
	chann = ch.NotifyPublish(chann)
	wg.Add(1)
	go func() {
		defer wg.Done()
		noti, ok := <-chann
		if !ok {
			return
		}
		if noti.Ack {
			rsp := PublishResponse{
				Code: 200,
				Body: []byte("Publish sucessful"),
			}
			log.Println("ack")
			result <- rsp
			return
		}
	}()
	for {
		res, ok := <-result
		log.Println("[RES]: ", res)
		if !ok {
			break
		}
		return &res, nil
	}
	wg.Wait()
	return nil, errors.New("Temp")
}
