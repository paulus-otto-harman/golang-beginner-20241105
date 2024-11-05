package middleware

import (
	"20241105/class/2/model"
	"20241105/class/2/repository"
	"20241105/database"
	"encoding/json"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database.DbOpen()
		defer db.Close()

		if err := repository.InitAuthRepo(db).Authorize(model.Session{Id: r.Header.Get("token")}); err != nil {
			json.NewEncoder(w).Encode(model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
