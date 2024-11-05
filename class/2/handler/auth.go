package handler

import (
	"20241105/class/2/model"
	"20241105/class/2/service"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func InitAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{AuthService: authService}
}

func (handle *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	json.NewDecoder(r.Body).Decode(&user)

	result := handle.AuthService.Login(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}
