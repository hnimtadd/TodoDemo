package business

import "todosService/internal/model"

type TodosRepository interface {
	AddTodo(req *AddTodoRequest) (*AddTodoResponse, error)
	GetTodos(req *GetTodosRequest) (*GetTodoResponse, error)
	UpdateTodo(id, field, value string) (*model.Todo, error)
	GetTodo(id string) (*model.Todo, error)
	DeleteTodo(req *DeleteTodoRequest) (*DeleteTodoResponse, error)
}

type AddTodoRequest struct {
	model.Todo
}

type AddTodoResponse struct {
	model.Todo
}

type GetTodosRequest struct {
	Size   int `json:"size"`
	Offset int `json:"offset"`
}
type GetTodoResponse struct {
	Result []model.Todo
}

type DeleteTodoRequest struct {
	Id string `json:"id"`
}

type DeleteTodoResponse struct {
	Result string `json:"result"`
}
