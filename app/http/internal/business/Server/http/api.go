package http

import (
	"encoding/json"
	publisher "httpService/internal/business/Server/http/Publisher"
	"httpService/internal/model"
	"httpService/protoc"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AddTodoRequest struct {
	Content  string         `json:"content"`
	Metadata model.Metadata `json:"metadata"`
}

type AddTodoResponse struct {
	Id          string         `json:"id"`
	Content     string         `json:"content"`
	Last_update time.Time      `json:"last_update"`
	Metadata    model.Metadata `json:"metadata"`
}

func (sv *FiberSever) AddTodoHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req AddTodoRequest
		if err := json.Unmarshal(ctx.Body(), &req); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		message := protoc.AddTodoRequest{
			Content: req.Content,
			Source:  req.Metadata.Source,
			Size:    strconv.Itoa(req.Metadata.Size),
			Format:  req.Metadata.Format,
		}

		body, err := json.Marshal(&message)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		log.Println("checkpoint")

		request := publisher.PublishRequest{
			ContentType: string(ctx.Request().Header.ContentType()),
			Action:      "insert",
			Body:        body,
		}

		response, err := sv.publisher.Publish(request)
		log.Println("checkpoint")
		if err != nil {
			log.Printf("error: %s", err.Error())
			ctx.SendStatus(fiber.ErrInternalServerError.Code)
			return ctx.SendString(err.Error())
		}

		// var responseMessage protoc.AddTodoResponse
		//
		// if err := json.Unmarshal(response.Body, &responseMessage); err != nil {
		// 	log.Panicf("error: %v", err)
		// 	json.NewEncoder(ctx.Response().BodyWriter()).Encode(map[string]string{"error": err.Error()})
		// 	return ctx.SendStatus(fiber.StatusInternalServerError)
		// }
		// log.Println("checkpoint")
		//
		// todo := responseMessage.GetResult()
		// Last_update, err := time.Parse(time.ANSIC, todo.GetLastUpdate())
		// if err != nil {
		// 	log.Printf("error: %s", err.Error())
		// 	ctx.SendStatus(fiber.ErrInternalServerError.Code)
		// 	return ctx.SendString("error")
		// }
		// log.Println("checkpoint")
		//
		// size, err := strconv.Atoi(todo.GetMetadata().GetSize())
		// if err != nil {
		// 	log.Printf("error: %s", err.Error())
		// 	ctx.SendStatus(fiber.ErrInternalServerError.Code)
		// 	return ctx.SendString("error")
		// }
		// log.Println("checkpoint")
		//
		// rsp := AddTodoResponse{
		// 	Id:          todo.GetId(),
		// 	Content:     todo.GetContent(),
		// 	Last_update: Last_update,
		// 	Metadata: model.Metadata{
		// 		Size:   size,
		// 		Format: todo.GetMetadata().GetFormat(),
		// 		Source: todo.GetMetadata().GetSource(),
		// 	},
		// }
		//
		// ctx.Response().Header.Set("Content-Type", "application/json")
		json.NewEncoder(ctx.Response().BodyWriter()).Encode(string(response.Body))
		return ctx.SendStatus(response.Code)
	}
}

type GetTodosRequest struct {
	Size   int `json:"size"`
	Offset int `json:"offset"`
}

type GetTodosResponse []*protoc.Todo

func (sv *FiberSever) GetTodosHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req GetTodosRequest
		if err := json.Unmarshal(ctx.Body(), &req); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		log.Printf("Get todos with size = %d, Offset = %d", req.Size, req.Offset)

		pReq := protoc.GetTodosRequest{
			Size:   int32(req.Size),
			Offset: int32(req.Offset),
		}

		body, err := json.Marshal(&pReq)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		request := publisher.PublishRequest{
			ContentType: string(ctx.Request().Header.ContentType()),
			Action:      "get",
			Body:        body,
		}
		response, err := sv.publisher.Publish(request)
		if err != nil {
			log.Printf("error: %v", err)
			json.NewEncoder(ctx.Response().BodyWriter()).Encode(map[string]string{"error": err.Error()})
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		var pRsp protoc.GetTodosResponse
		if err := json.Unmarshal(response.Body, &pRsp); err != nil {
			log.Printf("error: %v", err)
			json.NewEncoder(ctx.Response().BodyWriter()).Encode(map[string]string{"error": err.Error()})
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		rsp := pRsp.GetResults()

		ctx.Response().Header.Set("Content-Type", "application/json")
		json.NewEncoder(ctx.Response().BodyWriter()).Encode(&rsp)
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (sv *FiberSever) DeleteTodoHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		strid := ctx.Params("id")
		req := protoc.DeleteTodoRequest{
			Id: strid,
		}
		log.Println(strid)

		body, err := json.Marshal(&req)
		request := publisher.PublishRequest{
			ContentType: string(ctx.Request().Header.ContentType()),
			Action:      "delete",
			Body:        body,
		}
		response, err := sv.publisher.Publish(request)
		if err != nil {
			log.Printf("error: %v", err)
			json.NewEncoder(ctx.Response().BodyWriter()).Encode(map[string]string{"error": err.Error()})
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		var pRsp protoc.DeleteTodoResponse
		if err := json.Unmarshal(response.Body, &pRsp); err != nil {
			log.Printf("error: %v", err)
			json.NewEncoder(ctx.Response().BodyWriter()).Encode(map[string]string{"error": err.Error()})
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		log.Println(&pRsp)
		json.NewEncoder(ctx.Response().BodyWriter()).Encode(&pRsp)
		return ctx.SendStatus(fiber.StatusOK)
	}
}
