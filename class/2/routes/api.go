package route

import (
	"20241105/class/2/handler"
	"20241105/class/2/middleware"
	"20241105/class/2/repository"
	"20241105/class/2/service"
	"database/sql"
	"net/http"
)

// PublicApi Public API Routes
func PublicApi(db *sql.DB) *http.ServeMux {
	handleAuth := handler.InitAuthHandler(*service.InitAuthService(*repository.InitAuthRepo(db)))
	handleUser := handler.InitUserHandler(*service.InitUserService(*repository.InitUserRepo(db)))

	publicMux := http.NewServeMux()
	publicMux.HandleFunc("POST /register", handleUser.Registration)
	publicMux.HandleFunc("POST /login", handleAuth.Login)
	return publicMux
}

// ProtectedApi Protected API Routes
func ProtectedApi(db *sql.DB) http.Handler {
	handleUser := handler.InitUserHandler(*service.InitUserService(*repository.InitUserRepo(db)))
	handleTodo := handler.InitTodoHandler(*service.InitTodoService(*repository.InitTodoRepo(db)))

	userRoute := http.NewServeMux()

	userRoute.HandleFunc("GET /users", handleUser.All)
	userRoute.HandleFunc("GET /users/{id}", handleUser.Get)
	userRoute.HandleFunc("GET /todos", handleTodo.Get)
	userRoute.HandleFunc("POST /todos", handleTodo.Create)
	userRoute.HandleFunc("PUT /todos/{id}", handleTodo.Update)
	userRoute.HandleFunc("DELETE /todos/{id}", handleTodo.Delete)

	return middleware.Auth(userRoute)
}
