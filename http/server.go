package http

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vinibgoulart/sqleasy/helpers"
	databasesHandlers "github.com/vinibgoulart/sqleasy/packages/databases/handlers"
	"github.com/vinibgoulart/sqleasy/packages/server"
)

var state server.ServerState

func ServerInit(ctx context.Context, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	logger := helpers.LoggerCreate("HTTP Server")
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(JsonContentTypeMiddleware)
	router.Use(middleware.Timeout(30 * time.Second))

	router.Get("/status", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("OK"))
	})
	router.Post("/database-connect", databasesHandlers.DatabaseConnectPost(&state))

	router.Route("/database", func(r chi.Router) {
		r.Use(DbContextMiddleware(&state))
		r.Get("/info", databasesHandlers.DatabaseConnectInfoGet(&state))
		r.Get("/describe", databasesHandlers.DatabaseConnectDescribeGet(&state))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		logger.Info("HTTP server started on port 8080")
		server.ListenAndServe()
	}()

	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatalf("HTTP server shutdown error: %s", err)
	}
}
