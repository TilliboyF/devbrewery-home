package main

import (
	"log/slog"
	"net/http"

	"github.com/TilliboyF/devbrewery-home/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); nil != err {
		slog.Error("Can't load .env", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.Get("/", handler.MakeHandler(handler.HandleGetHome))

	slog.Info("Server running on port 8000")

	http.ListenAndServe(":8000", r)
}
