package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Movie struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Tagline string `json:"tagline"`
	Director string `json:"director"`
}

var MovieRepo = map[string]Movie {
	"1234": {Id: "1234", Title: "Joker", Tagline: "Put on a happy face.", Director: "Todd Phillips"},
	"5678": {Id: "5678", Title: "Fight Club", Tagline: "Mischief. Mayhem. Soap.", Director: "David Fincher"},
}

func AllMovies() []Movie {
	movies := make([]Movie, 0, len(MovieRepo))
	for _, v := range MovieRepo {
		movies = append(movies, v)
	}
	return movies
}

func MoviesHandleFunc(w http.ResponseWriter, r *http.Request)  {
	switch method := r.Method; method {
	case http.MethodGet:
		body, err := json.Marshal(AllMovies())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", "application/json; charset-utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeProblemResponse(w, http.StatusBadRequest, "Invalid body")
		}
		movie, err := FromJson(body)
		if err != nil {
			writeProblemResponse(w, http.StatusBadRequest, "Invalid body")
		}
		id, created := createMovie(movie)
		if created {
			w.Header().Add("Location", "/api/movies/"+id)
			w.WriteHeader(http.StatusCreated)
		}else{
			writeProblemResponse(w, http.StatusConflict, "Movie already exists")
		}
	default:
		writeProblemResponse(w, http.StatusBadRequest, "Unsupported request method")
	}
}

func MovieHandleFunc(w http.ResponseWriter, r *http.Request){
	id := r.URL.Path[len("/api/movies/"):]
	switch method := r.Method; method {
	case http.MethodGet:
		if movie, exists := MovieRepo[id]; exists {
			body, err := json.Marshal(movie)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Add("Content-Type", "application/json; charset-utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write(body)
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeProblemResponse(w, http.StatusBadRequest, "Invalid body")
		}
		movie, err := FromJson(body)
		if err != nil {
			writeProblemResponse(w, http.StatusBadRequest, "Invalid body")
		}
		exists := updateMovie(id, movie)
		if exists {
			w.WriteHeader(http.StatusNoContent)
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		existed := deleteMovie(id)
		if existed {
			w.WriteHeader(http.StatusNoContent)
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		writeProblemResponse(w, http.StatusBadRequest, "Unsupported request method")
	}
}

func createMovie(movie Movie) (string, bool) {
	if _, exists := MovieRepo[movie.Id]; exists {
		return "", false
	}else{
		MovieRepo[movie.Id] = movie
		return movie.Id, true
	}
}

func updateMovie(id string, movie Movie) (bool) {
	if _, exists := MovieRepo[id]; exists {
		movie.Id = id
		MovieRepo[id] = movie
		return true
	}else{
		return false
	}
}

func deleteMovie(id string) bool {
	if _, exists := MovieRepo[id]; exists {
		delete(MovieRepo, id)
		return true
	}else{
		return false
	}
}

func (m Movie) ToJson() []byte  {
	result, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return result
}

func FromJson(data []byte) (Movie, error) {
	movie := Movie{}
	err := json.Unmarshal(data, &movie)
	return movie, err
}

func writeProblemResponse(w http.ResponseWriter, status int, msg string) {
	w.Header().Add("Content-Type", "application/problem+json; charset-utf-8")
	w.WriteHeader(status)
	w.Write([]byte (`{"message":"`+msg+`"}`))
}