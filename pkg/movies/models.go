package movies

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title    string
	Year     string
	Genre    string
	ImdbID   string
	TmdbID   int
	Director string
	Actors   string //as CSV
	Runtime  string
	Released string
	Poster   string
	OverView string
}

// responses from movies APIs

// OmdbResponse is the response from http://www.omdbapi.com
// API call for search => http://www.omdbapi.com/?s={title}&apikey={key}
// API call for one movie => http://www.omdbapi.com/?i={IMDB_ID}&apikey={key}
// API call for one movie => http://www.omdbapi.com/?t={title}&apikey={key} will return most valued in IMDB
type OmdbResponse struct {
	Search []OmdbResult `json:"Search"`
}

type OmdbResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

// TmdbResponse is the response from https://api.themoviedb.org
// API call for search => https://api.themoviedb.org/3/search/movie?api_key={key}&language={lang}&query={title}
// API call for one movie => https://api.themoviedb.org/3/movie/{movieID}?api_key={key}&language={lang}
type TmdbResponse struct {
	Results []TmdbResult `json:"results"`
}

type TmdbResult struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Poster           string `json:"backdrop_path"`
	OriginalLanguage string `json:"original_language"`
	OriginalTitle    string `json:"original_title"`
	Overview         string `json:"overview"`
	PosterPath       string `json:"poster_path"`
	ReleaseDate      string `json:"release_date"`
}

type MovieSearchRes struct {
	Title  string `json:"title"`
	Poster string `json:"poster"`
	ImdbID string `json:"imdbID"`
	Year   string `json:"year"`
	Type   string `json:"type"`
	//to movie db api ID
	MovieDbId uint `json:"the_movie_db_id"`
}

func (o *OmdbResponse) ToApi() []MovieSearchRes {
	var res []MovieSearchRes
	for _, s := range o.Search {
		res = append(res, MovieSearchRes{
			Title:  s.Title,
			Poster: s.Poster,
			ImdbID: s.ImdbID,
			Year:   s.Year,
			Type:   s.Type,
		})
	}
	return res
}

//UpcomingMovies response from API
type UpcomingMovies struct {
	Results []UpcomingResult `json:"results"`
}

type UpcomingResult struct {
	ID               uint   `json:"id"`
	Adult            bool   `json:"adult"`
	BackdropPath     string `json:"backdrop_path"`
	OriginalLanguage string `json:"original_language"`
	OriginalTitle    string `json:"original_title"`
	Overview         string `json:"overview"`
	PosterPath       string `json:"poster_path"`
	ReleaseDate      string `json:"release_date"`
	Title            string `json:"title"`
}

func (u *UpcomingMovies) ToApi() []MovieSearchRes {
	var res []MovieSearchRes

	for _, s := range u.Results {
		res = append(res, MovieSearchRes{
			Title:  s.Title,
			Poster: s.PosterPath,
			Year:      s.ReleaseDate,
			MovieDbId: s.ID,
		})
	}
	return res
}
