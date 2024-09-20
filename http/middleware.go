package http

import (
	"net/http"

	"github.com/vinibgoulart/sqleasy/helpers"
	"github.com/vinibgoulart/sqleasy/packages/server"
)

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func DbContextMiddleware(state *server.ServerState) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			if state.Db == nil {
				helpers.ErrorResponse("Database not found, connect in POST /connect-database", http.StatusBadRequest, res)
				return
			}

			errPing := state.Db.Ping()

			if errPing != nil {
				helpers.ErrorResponse("Database ping not working, check your connection", http.StatusBadRequest, res)
				return
			}

			if state.DatabaseConnect == nil {
				helpers.ErrorResponse("Database connection not found, connect in POST /connect-database", http.StatusBadRequest, res)
				return
			}

			next.ServeHTTP(res, req)
		})
	}
}
