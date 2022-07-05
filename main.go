package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

func (m *Movie) Validation() error {
	if m.Isbn == "" {
		return errors.New("isbn field empty")
	}
	if m.Title == "" {
		return errors.New("title field empty")
	}
	if m.Director.FirstName == "" {
		return errors.New("first name field empty")
	}
	if m.Director.LastName == "" {
		return errors.New("last name field empty")
	}
	return nil
}

type Director struct {
	FirstName string `josn:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	if err := movie.Validation(); err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			if err := movie.Validation(); err != nil {
				fmt.Fprint(w, err.Error())
				return
			}
			movie.ID = params["id"]
			movies[index] = movie
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "438227",
		Title: "Movie One",
		Director: &Director{
			FirstName: "John",
			LastName:  "Doe",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "45455",
		Title: "Movie Two",
		Director: &Director{
			FirstName: "Steve",
			LastName:  "Smith",
		},
	})
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Println("Starting server http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
