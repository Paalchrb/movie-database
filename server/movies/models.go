package movies

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Paalchrb/movie-database/server/config"
)

//Movie struct represent a movie instance
type Movie struct {
	ID       int
	Title    string
	Poster   string
	Year     string
	Genre    string
	Rating   string
	Duration string
	Summary  string
}

//GetAllMovies returns all rows from movie table
func GetAllMovies() ([]Movie, error) {
	rows, err := config.DB.Query("SELECT * FROM movies ORDER BY rating DESC;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	movies := make([]Movie, 0)
	for rows.Next() {
		movie := Movie{}
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Poster, &movie.Year, &movie.Genre, &movie.Rating, &movie.Duration, &movie.Summary) // order matters
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}

//GetMovieByID returns one row from movie table, based on provided id
func GetMovieByID(req *http.Request) (Movie, error) {
	movie := Movie{}
	id := req.FormValue("id")
	if id == "" {
		return movie, errors.New("400. Bad request")
	}

	row := config.DB.QueryRow("SELECT * FROM movies WHERE id=$1;", id)

	err := row.Scan(&movie.ID, &movie.Title, &movie.Poster, &movie.Year, &movie.Genre, &movie.Rating, &movie.Duration, &movie.Summary) // order matters
	if err != nil {
		return movie, err
	}

	return movie, nil
}

//CreateMovie get formdata from template and insert new row in movies table
func CreateMovie(req *http.Request) (Movie, error) {
	movie := Movie{}
	movie.Title = req.FormValue("title")
	movie.Poster = req.FormValue("poster")
	movie.Year = req.FormValue("year")
	movie.Genre = req.FormValue("genre")
	movie.Rating = req.FormValue("rating")
	movie.Duration = req.FormValue("duration")
	movie.Summary = req.FormValue("summary")

	// validate form values
	if movie.Title == "" ||
		movie.Poster == "" ||
		movie.Year == "" ||
		movie.Genre == "" ||
		movie.Rating == "" ||
		movie.Duration == "" ||
		movie.Summary == "" {
		return movie, errors.New("400. Bad request. All fields must be complete")
	}

	//insert values
	_, err := config.DB.Exec("INSERT INTO movies (title, poster, year, genre, rating, duration, summary) VALUES ($1, $2, $3, $4, $5, $6, $7);", movie.Title, movie.Poster, movie.Year, movie.Genre, movie.Rating, movie.Duration, movie.Summary)
	if err != nil {
		return movie, errors.New("500. Internal Server Error." + err.Error())
	}

	return movie, nil
}

//UpdateMovie get formdata from template and updates row with provided id
func UpdateMovie(req *http.Request, i string) (Movie, error) {
	movie := Movie{}
	movie.Title = req.FormValue("title")
	movie.Poster = req.FormValue("poster")
	movie.Year = req.FormValue("year")
	movie.Genre = req.FormValue("genre")
	movie.Rating = req.FormValue("rating")
	movie.Duration = req.FormValue("duration")
	movie.Summary = req.FormValue("summary")
	id, _ := strconv.Atoi(i)

	// validate form values
	if movie.Title == "" ||
		movie.Poster == "" ||
		movie.Year == "" ||
		movie.Genre == "" ||
		movie.Rating == "" ||
		movie.Duration == "" ||
		movie.Summary == "" {
		return movie, errors.New("400. Bad request. All fields must be complete")
	}

	//insert values
	_, err := config.DB.Exec("UPDATE movies SET title = $1, poster = $2, year = $3, genre = $4, rating = $5, duration = $6, summary = $7 WHERE id = $8;", movie.Title, movie.Poster, movie.Year, movie.Genre, movie.Rating, movie.Duration, movie.Summary, id)
	if err != nil {
		fmt.Println(err)
		return movie, errors.New("500. Internal Server Error." + err.Error())
	}

	return movie, nil
}

//DeleteBook removes the row with corresponding ID from database
func DeleteBook(req *http.Request) error {
	id := req.FormValue("id")
	if id == "" {
		return errors.New("400. Bad request")
	}

	_, err := config.DB.Exec("DELETE FROM movies WHERE id=$1;", id)
	if err != nil {
		return errors.New("500. Internal server error")
	}
	return nil
}
