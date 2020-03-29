package movies

import (
	"net/http"

	"github.com/Paalchrb/movie-database/server/config"
)

//Index fetches data about all movies and loads template
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	movies, err := AllMovies()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "movies.gohtml", movies)
}
