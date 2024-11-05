package service

import (
	"20241105/class/2/repository"
	"net/http"
)

type WebPageService struct {
	WebPage repository.WebPage
}

func InitWebPageService(repo repository.WebPage) *WebPageService {
	return &WebPageService{WebPage: repo}
}

func (repo *WebPageService) Render(w http.ResponseWriter, templateName string, title string) {
	repo.WebPage.Render(w, templateName, title)
}
