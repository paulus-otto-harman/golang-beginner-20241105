package service

import (
	"20241105/class/2/model"
	"20241105/class/2/repository"
)

type AuthService struct {
	Auth repository.Auth
}

func InitAuthService(repo repository.Auth) *AuthService {
	return &AuthService{Auth: repo}
}

func (repo *AuthService) Login(user model.User) *model.Response {
	session, err := repo.Auth.Authenticate(&user)
	if err != nil {
		return &model.Response{StatusCode: 401, Message: "Authentication failed", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "User authenticated, Session created", Data: session}
}

func (repo *AuthService) Logout(session model.Session) *model.Response {
	s, err := repo.Auth.Logout(&session)
	if err != nil {
		return &model.Response{StatusCode: 401, Message: "Logout Failed", Data: err.Error()}
	}
	if s == 0 {
		return &model.Response{StatusCode: 401, Message: "Invalid user session", Data: nil}
	}
	return &model.Response{StatusCode: 200, Message: "User successfully logged out", Data: nil}
}
