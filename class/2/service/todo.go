package service

import (
	"20241105/class/2/model"
	"20241105/class/2/repository"
)

type TodoService struct {
	Todo repository.Todo
}

func InitTodoService(repo repository.Todo) *TodoService {
	return &TodoService{Todo: repo}
}

func (repo *TodoService) Get(session model.Session) *model.Response {
	todos, err := repo.Todo.Get(session)
	if err != nil {
		return &model.Response{StatusCode: 500, Message: "Server Error", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "Successfully get all Todos", Data: todos}
}

func (repo *TodoService) Create(todo model.Todo, session model.Session) *model.Response {
	err := repo.Todo.Create(&todo, session)
	if err != nil {
		return &model.Response{StatusCode: 500, Message: "Server Error", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "Todo created", Data: todo}
}

func (repo *TodoService) Update(todo model.Todo, session model.Session) *model.Response {
	err := repo.Todo.Update(&todo, session)
	if err != nil {
		return &model.Response{StatusCode: 404, Message: "Todo Not Found", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "Todo updated", Data: todo}
}

func (repo *TodoService) Delete(todo model.Todo, session model.Session) *model.Response {
	err := repo.Todo.Delete(&todo, session)
	if err != nil {
		return &model.Response{StatusCode: 404, Message: "Todo Not Found", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "Todo deleted", Data: todo}
}
