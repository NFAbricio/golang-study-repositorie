package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
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
	UpdateMovie(ctx context.Context, movie *Movie) (*Movie, error)
	DeleteMovie(ctx context.Context, id string) error
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


func (c redisCache) GetMovie(ctx context.Context, id string) (*Movie, error) {
	logger := zap.Must(zap.NewProduction())
	cache := c.getClient()
	val, err := cache.HGet(ctx, "movies", id).Result()
	if err != nil {
		logger.Error("error to get movie", zap.Error(errors.New("something happened")))
		return nil, err
	}

	movie := &Movie{}
	err = json.Unmarshal([]byte(val), movie)
	if err != nil {
		logger.Error("error to unmarshal movie", zap.Error(errors.New("something happened")))
		return nil, err
	}
	
	return movie, nil
}

func (c redisCache) GetMovies(ctx context.Context) ([]*Movie, error) {
	logger := zap.Must(zap.NewProduction())
	cache := c.getClient()
	movies := []*Movie{}

	val, err:= cache.HGetAll(ctx, "movies").Result()
	if err != nil {
		logger.Error("error to get movies", zap.Error(errors.New("something happened")))
		return nil, err
	}

	for _, item := range val{
		movie := &Movie{}
		err := json.Unmarshal([]byte(item), movie)
		if err != nil {
			logger.Error("error to Unmarshal movie", zap.Error(errors.New("something happened")))
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}


func (c redisCache) CreateMovie(ctx context.Context, movie *Movie) (*Movie, error) {
	logger := zap.Must(zap.NewProduction())
	cache := c.getClient()

	movie.Id = uuid.New().String()

	json, err := json.Marshal(movie)//?
	if err != nil {
		logger.Error("error to marshal movie", zap.Error(errors.New("something happened")))
		return nil, err
	}

	cache.HSet(ctx, "movies", movie.Id, json)//?
	if err != nil{
		logger.Error("error to set movie", zap.Error(errors.New("something happened")))
		return nil, err
	}

	return movie, nil
}

func (c redisCache) UpdateMovie(ctx context.Context,movie *Movie) (*Movie, error) {
	logger := zap.Must(zap.NewProduction())
	cache := c.getClient()

	json, err := json.Marshal(movie)
	if err != nil {
		logger.Error("error to marshal movie", zap.Error(errors.New("something happened")))
		return nil, err
	}

	cache.HSet(ctx, "movies", movie.Id, json)

	return movie, nil
}


func (c redisCache) DeleteMovie(ctx context.Context,id string) error {
	logger := zap.Must(zap.NewProduction())
	cache := c.getClient()

	numDeleted, err := cache.HDel(ctx,"movies", id).Result() //if the movie is deleted, numDeleted receive 1
	if numDeleted == 0 {
		return errors.New("movie to delete not found")
	}
	if err != nil {
		logger.Error("error to delete movie", zap.Error(errors.New("something happened")))
		return err
	}

	return nil
}