package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-uuid"
	"github.com/redis/go-redis/v9"
)

type Movie struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MovieService interface {
	GetMovie(ctx context.Context, id string) (*Movie, error)
	GetMovies(ctx context.Context) ([]*Movie, error)
	CreateMovie(ctx context.Context, movie *Movie) (*Movie, error)
	UpdateMovie(movie *Movie) (*Movie, error)
	DeleteMovie(id string) error
}

type redisCache struct {
	host string
	db   int
	exp  time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) MovieService {
	return &redisCache{
		host: host,
		db: db,
		exp: exp,
	}
}

func (c redisCache) getClient() *redis.Client{
	return redis.NewClient(&redis.Options{
		Addr: c.host,
		Password: "",
		DB: c.db,
	})
}

//add zap logger after
func (c redisCache) GetMovie(ctx context.Context, id string) (*Movie, error) {
	cache := c.getClient()
	val, err := cache.HGet(ctx, "movies", id).Result()
	if err != nil {
		return nil, err
	}

	movie := &Movie{}
	err = json.Unmarshal([]byte(val), movie)
	if err != nil {
		return nil, err
	}
	
	return movie, nil
}

func (c redisCache) GetMovies(ctx context.Context) ([]*Movie, error) {
	cache := c.getClient()
	movies := []*Movie{}

	val, err:= cache.HGetAll(ctx, "movies").Result()
	if err != nil {
		return nil, err
	}

	for _, item := range val{
		movie := &Movie{}
		err := json.Unmarshal([]byte(item), movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}


func (c redisCache) CreateMovie(ctx context.Context, movie *Movie) (*Movie, error) {
	cache := c.getClient()

	movie.Id = uuid.New().String()

	json, err := json.Marshal(movie)//?
	if err != nil {
		return nil, err
	}

	cache.HSet(ctx, "movies", movie.Id, json)//?
	if err != nil{
		return nil, err
	}

	return movie, nil
}

func (c redisCache) UpdateMovie(movie *Movie) (*Movie, error) {

}

func (c redisCache) DeleteMovie(id string) error {

}