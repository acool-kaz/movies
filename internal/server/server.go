package server

import (
	"log"
	"net/http"

	"github.com/acool-kaz/movies/internal/handler"
	"github.com/acool-kaz/movies/models"
)

func Run() error {
	handler := handler.NewHandler()
	handler.InitEndpoits()

	handler.Movies = append(handler.Movies, models.Movie{
		ID:    "1",
		Isbn:  "438227",
		Title: "Movie One",
		Director: &models.Director{
			FirstName: "John",
			LastName:  "Doe",
		},
	})
	handler.Movies = append(handler.Movies, models.Movie{
		ID:    "2",
		Isbn:  "45455",
		Title: "Movie Two",
		Director: &models.Director{
			FirstName: "Steve",
			LastName:  "Smith",
		},
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: handler.Router,
	}

	log.Println("Starting server http://localhost:8080")
	return server.ListenAndServe()
}
