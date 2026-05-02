package server

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHTTPServer() *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		// back to using this stuff
		port = "8080"
	}
	addr := ":" + port
	httpServer := newHTTPServer()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", httpServer.handleConsume)
	r.Post("/", httpServer.handleProduce)
	return &http.Server{
		Addr: addr,
		Handler: r,

		ReadTimeout: 5 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}
}
