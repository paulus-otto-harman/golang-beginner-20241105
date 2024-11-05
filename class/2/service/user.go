package service

import (
	"20241105/class/2/model"
	"20241105/class/2/repository"
	"log"
)

type UserService struct {
	User repository.User
}

func InitUserService(repo repository.User) *UserService {
	return &UserService{User: repo}
}

func (repo *UserService) Create(user model.User) *model.Response {
	session, err := repo.User.Create(&user)
	if err != nil {
		return &model.Response{StatusCode: 500, Message: "Server Error", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "User successfully registered and login session created", Data: session}
}

func (repo *UserService) All(session model.Session) *model.Response {
	users, err := repo.User.All(session)
	if err != nil {
		return &model.Response{StatusCode: 500, Message: "Server Error", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "Successfully get all Users", Data: users}
}

func (repo *UserService) Get(user model.User) *model.Response {
	session := model.Session{User: user}
	err := repo.User.Get(&session)
	log.Println(session)
	if err != nil {
		return &model.Response{StatusCode: 500, Message: "Server Error", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "User successfully retrieved", Data: session}
}
