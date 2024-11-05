package main

import (
	route "20241105/class/2/routes"
	"20241105/database"
	"log"
	"net/http"
)

func main() {
	db := database.DbOpen()
	defer db.Close()

	serverMux := http.NewServeMux()
	serverMux.Handle("/", route.WebTemplate())
	serverMux.Handle("/api/", http.StripPrefix("/api", route.PublicApi(db)))
	serverMux.Handle("/api/user/", http.StripPrefix("/api/user", route.ProtectedApi(db)))

	log.Println("Server started on port 8080")

	http.ListenAndServe(":8080", serverMux)

}
