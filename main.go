package main

import (
	"log/slog"
	"net/http"

	"github.com/TilliboyF/devbrewery-home/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	r.Get("/", handler.MakeHandler(handler.HandleGetHome))

	slog.Info("Server running on port 8000")

	http.ListenAndServe(":8000", r)
}
