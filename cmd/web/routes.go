package main

import (
	"net/http"

	"github.com/flaviusp23/bookings/pkg/config"
	"github.com/flaviusp23/bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// http.Handle("/tmpfiles/",
	// http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
	// Explanation:

	// FileServer() is told the root of static files is "/tmp". We want the URL to start with "/tmpfiles/". So if someone requests "/tempfiles/example.txt", we want the server to send the file "/tmp/example.txt".

	// In order to achieve this, we have to strip "/tmpfiles" from the URL, and the remaining will be the relative path compared to the root folder "/tmp" which if we join gives:

	// /tmp/example.txt
	return mux
}
