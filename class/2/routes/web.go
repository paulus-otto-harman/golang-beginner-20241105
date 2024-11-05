package route

import (
	"20241105/class/2/handler"
	"20241105/class/2/repository"
	"20241105/class/2/service"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
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

	webMux.HandleFunc("/register.js", staticHandler)
	webMux.HandleFunc("/users.js", staticHandler)
	webMux.HandleFunc("/todos.js", staticHandler)
	webMux.HandleFunc("/app.css", staticHandler)

	return webMux
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Println(path)
	var data []byte

	if strings.HasSuffix(path, "js") {
		data, _ = os.ReadFile(fmt.Sprintf("templates/%s", path))
		w.Header().Set("Content-Type", "application/javascript")
	} else {
		data, _ = os.ReadFile("templates/app.css")
		w.Header().Set("Content-Type", "text/css")
	}
	w.Write(data)
}
