package movies

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/pelix-api/pkg/users"
	"gorm.io/gorm"
	"net/http"
	"os"
	"sync"
)

type MovieStorage interface {
	Lookup(title string) ([]MovieSearchRes, error)
	Upcoming() ([]MovieSearchRes, error)
	//SaveLocal save in local DB
	SaveLocal(movies []MovieSearchRes)
}

type ListStorage interface {
	Add(item ListItem) error
	Update(itemId uint) error
}

func (l *listStorage) Add(item ListItem) error {
	//check if the list is created
	var list WatchList
	rows := l.db.Where("user_id=?", item.UserId).First(&list).RowsAffected

	if rows == 0 {
		l.db.Create(&list)
	}

	var user users.User
	err := l.db.Where("id=?", item.UserId).First(&user).Error

	if err != nil {
		return err
	}

	var movie Movie
	err = l.db.Where("id=?", item.MovieId).First(&movie).Error

	if err != nil {
		return err
	}

	listItem := WatchList{
		User:   user,
		Movies: append(list.Movies, movie),
		Seen:   false,
	}

	return l.db.Save(&listItem).Error
}

func (l *listStorage) Update(itemId uint) error {
	panic("implement me")
}

type movieStorage struct {
	db *gorm.DB
}

type listStorage struct {
	db *gorm.DB
}

func newMovieStorage(db *gorm.DB) *movieStorage {
	return &movieStorage{
		db: db,
	}
}

func newListStorage(db *gorm.DB) *listStorage {
	return &listStorage{
		db: db,
	}
}

func (m *movieStorage) Lookup(title string) ([]MovieSearchRes, error) {
	result, err := search(title)

	if err != nil {
		return nil, err
	}

	go m.SaveLocal(result)
	return result, nil
}

func (m *movieStorage) Upcoming() ([]MovieSearchRes, error) {
	result, err := upcoming()

	if err != nil {
		return nil, err
	}

	go m.SaveLocal(result)

	return result, err
}

func (m *movieStorage) SaveLocal(movies []MovieSearchRes) {
	var wg sync.WaitGroup
	for _, movie := range movies {
		wg.Add(1)
		go save(movie, m.db, &wg)
	}
	wg.Wait()
}

func save(m MovieSearchRes, db *gorm.DB, wg *sync.WaitGroup) {
	movie := m.ToMovie()

	err := db.Create(movie).Error

	if err != nil {
		log.Error().Err(err)
	}

	wg.Done()
}

const uriFmt = "https://www.omdbapi.com/?s=%s&apikey=4ecb0111"

func search(t string) ([]MovieSearchRes, error) {
	res, err := http.Get(fmt.Sprintf(uriFmt, t))

	if err != nil {
		return nil, err
	}

	body := res.Body
	defer func() {
		_ = body.Close()
	}()

	var omdbRes OmdbResponse
	err = json.NewDecoder(body).Decode(&omdbRes)

	if err != nil {
		return nil, err
	}

	return omdbRes.ToApi(), nil
}

func upcoming() ([]MovieSearchRes, error) {
	uriFmt := "https://api.themoviedb.org/3/movie/upcoming?api_key=%s"
	uri := fmt.Sprintf(uriFmt, os.Getenv("MOVIE_KEY"))

	res, err := http.Get(uri)

	if err != nil {
		return nil, err
	}

	var movies UpcomingMovies
	body := res.Body
	defer func() {
		_ = body.Close()
	}()
	err = json.NewDecoder(body).Decode(&movies)

	return movies.ToApi(), nil
}
