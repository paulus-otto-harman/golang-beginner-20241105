package route

import (
	"20241105/class/2/handler"
	"20241105/class/2/repository"
	"20241105/class/2/service"
	"html/template"
	"log"
	"net/http"
)

func initTemplate() (*repository.WebPageData, *template.Template) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
		return nil, nil
	}

	return &repository.WebPageData{}, tmpl
}

// Web Template Routes
func WebTemplate() *http.ServeMux {
	handleWebPage := handler.InitWebPageHandler(*service.InitWebPageService(*repository.InitWebPageRepo(initTemplate())))

	webMux := http.NewServeMux()
	webMux.HandleFunc("/", handleWebPage.Home)
	webMux.HandleFunc("/register", handleWebPage.Registration)
	webMux.HandleFunc("/users", handleWebPage.Users)
	webMux.HandleFunc("/todos", handleWebPage.Todos)

	return webMux
}
