package main

import (
	"20241105/class/2/handler"
	"20241105/class/2/middleware"
	"20241105/class/2/repository"
	"20241105/class/2/service"
	"20241105/database"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db := database.DbOpen()
	defer db.Close()

	handleAuth := handler.InitAuthHandler(*service.InitAuthService(*repository.InitAuthRepo(db)))
	handleUser := handler.InitUserHandler(*service.InitUserService(*repository.InitUserRepo(db)))
	handleTodo := handler.InitTodoHandler(*service.InitTodoService(*repository.InitTodoRepo(db)))
	handleWebPage := handler.InitWebPageHandler(*service.InitWebPageService(*repository.InitWebPageRepo(InitTemplate())))

	// Web Routes
	webMux := http.NewServeMux()
	webMux.HandleFunc("/", handleWebPage.Home)
	webMux.HandleFunc("/register", handleWebPage.Registration)
	webMux.HandleFunc("/users", handleWebPage.Users)
	webMux.HandleFunc("/todos", handleWebPage.Todos)

	// Public API Routes
	publicMux := http.NewServeMux()

	publicMux.HandleFunc("POST /register", handleUser.Registration)
	publicMux.HandleFunc("POST /login", handleAuth.Login)

	// Protected API Routes
	userRoute := http.NewServeMux()

	userRoute.HandleFunc("GET /users", handleUser.All)
	userRoute.HandleFunc("GET /users/{id}", handleUser.Get)
	userRoute.HandleFunc("GET /todos", handleTodo.Get)
	userRoute.HandleFunc("POST /todos", handleTodo.Create)
	userRoute.HandleFunc("PUT /todos/{id}", handleTodo.Update)
	userRoute.HandleFunc("DELETE /todos/{id}", handleTodo.Delete)

	protectedMux := middleware.Auth(userRoute)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", webMux)
	serverMux.Handle("/api/", http.StripPrefix("/api", publicMux))
	serverMux.Handle("/api/user/", http.StripPrefix("/api/user", protectedMux))

	log.Println("Server started on port 8080")

	http.ListenAndServe(":8080", serverMux)

}

func InitTemplate() (*repository.WebPageData, *template.Template) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
		return nil, nil
	}

	return &repository.WebPageData{}, tmpl
}
