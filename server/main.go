package main

import (
	"net/http"

	"github.com/Paalchrb/movie-database/server/movies"
)

func main() {
	//Handles browser equest to favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", index)
	http.HandleFunc("/movies", movies.Index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/movies", http.StatusSeeOther)
}
