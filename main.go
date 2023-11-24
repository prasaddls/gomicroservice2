package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "234456", Title: "RAGA", Director: &Director{Firstname: "Raja", Lastname: "Mouly"}})
	movies = append(movies, Movie{ID: "2", Isbn: "234456", Title: "RAGA", Director: &Director{Firstname: "Raja", Lastname: "Mouly"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	//r.HandleFunc("/movies", createMovie).Methods("POST")
	//r.HandleFunc("/movies/{id}", upateMovie).Method("PUT")
	//r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")

	fmt.Print(" Starting server at port 8080: \n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
