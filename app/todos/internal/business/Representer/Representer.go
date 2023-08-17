package representer

import (
	"strconv"
	"time"
	"todosService/internal/business"
	"todosService/internal/model"
	"todosService/protoc"

	"github.com/google/uuid"
)

type Representer interface {
	AddTodoTransform(protoc.AddTodoRequest) business.AddTodoRequest
	GetTodoTransfrom(protoc.GetTodosRequest) business.GetTodosRequest
}

func AddTodoTransform(req *protoc.AddTodoRequest) (*business.AddTodoRequest, error) {
	id := uuid.New().String()
	size, err := strconv.Atoi(req.Size)
	if err != nil {
		return nil, err
	}
	todo := business.AddTodoRequest{
		Todo: model.Todo{
			Id:          id,
			Last_update: time.Now(),
			Is_deleted:  false,
			Content:     req.Content,
			Metadata: model.Metadata{
				Source: req.Source,
				Size:   size,
				Format: req.Format,
			},
		}}
	return &todo, nil
}

func AddTodoDeTransfrom(reporsp *business.AddTodoResponse) (*protoc.AddTodoResponse, error) {
	size := strconv.Itoa(reporsp.Metadata.Size)
	rsp := protoc.AddTodoResponse{
		Result: &protoc.Todo{
			Id:         reporsp.Id,
			Content:    reporsp.Content,
			LastUpdate: reporsp.Last_update.Format(time.ANSIC),
			Metadata: &protoc.TodoMeta{
				Source: reporsp.Metadata.Source,
				Format: reporsp.Metadata.Format,
				Size:   size,
			},
		},
	}
	return &rsp, nil
}
func GetTodoTransfrom(req *protoc.GetTodosRequest) (*business.GetTodosRequest, error) {
	bRequest := business.GetTodosRequest{
		Size:   int(req.GetSize()),
		Offset: int(req.GetOffset()),
	}
	return &bRequest, nil
}
func GetTodoDeTransfrom(rsp *business.GetTodoResponse) (*protoc.GetTodosResponse, error) {
	results := []*protoc.Todo{}
	for _, bTodo := range rsp.Result {
		pTodo := protoc.Todo{
			Id:         bTodo.Id,
			Content:    bTodo.Content,
			LastUpdate: bTodo.Last_update.Format(time.ANSIC),
			Metadata: &protoc.TodoMeta{
				Source: bTodo.Metadata.Source,
				Size:   strconv.Itoa(bTodo.Metadata.Size),
				Format: bTodo.Metadata.Format,
			},
		}
		results = append(results, &pTodo)
	}
	pRsp := protoc.GetTodosResponse{
		Results: results,
	}
	return &pRsp, nil
}

func DeleteTodoTransform(req *protoc.DeleteTodoRequest) (*business.DeleteTodoRequest, error) {
	bReq := business.DeleteTodoRequest{
		Id: req.GetId(),
	}
	return &bReq, nil
}

func DeleteTodoDetransform(rsp *business.DeleteTodoResponse) (*protoc.DeleteTodoResponse, error) {
	pRsp := protoc.DeleteTodoResponse{
		Result: rsp.Result,
	}
	return &pRsp, nil
}
