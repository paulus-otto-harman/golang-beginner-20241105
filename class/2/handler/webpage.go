package handler

import (
	"20241105/class/2/service"
	"net/http"
)

type WebPageHandler struct {
	WebPageService service.WebPageService
}

func InitWebPageHandler(webPageService service.WebPageService) WebPageHandler {
	return WebPageHandler{WebPageService: webPageService}
}

func (handle *WebPageHandler) Home(w http.ResponseWriter, r *http.Request) {
	handle.WebPageService.Render(w, "index.html", "Home")
}

func (handle *WebPageHandler) Registration(w http.ResponseWriter, r *http.Request) {
	handle.WebPageService.Render(w, "register.html", "Registration")
}

func (handle *WebPageHandler) Users(w http.ResponseWriter, r *http.Request) {
	handle.WebPageService.Render(w, "users.html", "List Data User")
}

func (handle *WebPageHandler) Todos(w http.ResponseWriter, r *http.Request) {
	handle.WebPageService.Render(w, "todos.html", "Todo List")
}
