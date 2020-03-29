package movies

import (
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

//AllMovies fetches all rows from database
func AllMovies() ([]Movie, error) {
	rows, err := config.DB.Query("SELECT * FROM movies")
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
