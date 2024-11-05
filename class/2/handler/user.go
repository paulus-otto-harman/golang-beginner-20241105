package handler

import (
	"20241105/class/2/model"
	"20241105/class/2/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService service.UserService
}

func InitUserHandler(userService service.UserService) UserHandler {
	return UserHandler{UserService: userService}
}

func (handle *UserHandler) Registration(w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	json.NewDecoder(r.Body).Decode(&user)

	result := handle.UserService.Create(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}

func (handle *UserHandler) All(w http.ResponseWriter, r *http.Request) {
	result := handle.UserService.All(model.Session{Id: r.Header.Get("token")})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (handle *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	result := handle.UserService.Get(model.User{Id: id})
	log.Println(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
