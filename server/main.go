package main

import (
	"net/http"

	"github.com/Paalchrb/movie-database/server/movies"
)

func main() {
	//Handles browser equest to favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())

	//Serves style.css
	http.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./templates/styles"))))

	//Routing
	http.HandleFunc("/", index)
	http.HandleFunc("/movies", movies.ShowAll)
	http.HandleFunc("/movies/show", movies.ShowOne)
	http.HandleFunc("/movies/create", movies.Create)
	http.HandleFunc("/movies/update", movies.Update)

	//Listen and serve at port 3000:
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/movies", http.StatusSeeOther)
}
