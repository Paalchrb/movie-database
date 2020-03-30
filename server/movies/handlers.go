package movies

import (
	"database/sql"
	"net/http"

	"github.com/Paalchrb/movie-database/server/config"
)

//ShowAll fetches data about all movies and loads template
func ShowAll(w http.ResponseWriter, req *http.Request) {
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

//Create adds a new movie instance to database:
func Create(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
		break
	case "POST":
		createMovie(w, req)
	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}

func createMovie(w http.ResponseWriter, req *http.Request) {
	movie, err := CreateMovie(req)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "created.gohtml", movie)
}

//Update changes the movie with the provided ID:
func Update(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	switch req.Method {
	case "GET":
		renderForm(w, req)
	case "POST":
		updateMovie(w, req, id)
	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}

func renderForm(w http.ResponseWriter, req *http.Request) {
	movie, err := GetMovieByID(req)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, req)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "update.gohtml", movie)
}

func updateMovie(w http.ResponseWriter, req *http.Request, id string) {
	movie, err := UpdateMovie(req, id)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "created.gohtml", movie)
}

//Delete deletes the movie with the provided ID:
func Delete(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteBook(req)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}

	http.Redirect(w, req, "/books", http.StatusSeeOther)
}
