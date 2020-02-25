package main

import (
	"fmt"
	"go-movies/api"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/api/movies", api.MoviesHandleFunc)
	http.HandleFunc("/api/movies/", api.MovieHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("MOVIE_PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":"+port
}

func index(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Alive!")
}
