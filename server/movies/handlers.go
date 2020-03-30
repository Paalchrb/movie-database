package movies

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Paalchrb/movie-database/server/config"
)

//ShowAll fetches data about all movies and loads template
func ShowAll(w http.ResponseWriter, req *http.Request) {
	fmt.Println("New request:", req.URL, req.Method)
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	movies, err := GetAllMovies()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "movies.gohtml", movies)
}

//ShowOne fetches data about one movie, based on unique ID:
func ShowOne(w http.ResponseWriter, req *http.Request) {
	fmt.Println("New request:", req.URL, req.Method)
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	movie, err := GetMovieByID(req)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, req)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "movie.gohtml", movie)
}
