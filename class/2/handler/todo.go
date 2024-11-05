package handler

import (
	"20241105/class/2/model"
	"20241105/class/2/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	TodoService service.TodoService
}

func InitTodoHandler(todoService service.TodoService) TodoHandler {
	return TodoHandler{TodoService: todoService}
}

func (handler TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	todo := model.Todo{}

	json.NewDecoder(r.Body).Decode(&todo)

	result := handler.TodoService.Create(todo, model.Session{Id: r.Header.Get("token")})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}

func (handler TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	result := handler.TodoService.Get(model.Session{Id: r.Header.Get("token")})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (handler TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	result := handler.TodoService.Update(model.Todo{Id: id}, model.Session{Id: r.Header.Get("token")})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (handler TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	result := handler.TodoService.Delete(model.Todo{Id: id}, model.Session{Id: r.Header.Get("token")})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
