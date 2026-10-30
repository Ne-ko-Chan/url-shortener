package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
}

type Config struct {
	Addr        string        `yaml:"addr"`
	Password    string        `yaml:"password"`
	User        string        `yaml:"user"`
	DB          int           `yaml:"dB"`
	MaxRetries  int           `yaml:"max_retries"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
	Timeout     time.Duration `yaml:"timeout"`
}

func NewStorage(ctx context.Context, cfg Config) (*RedisStorage, error) {
	db := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Username:     cfg.User,
		Password:     cfg.Password,
		DB:           cfg.DB,
		MaxRetries:   cfg.MaxRetries,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %v", err)
	}

	return &RedisStorage{db}, nil
}

func (s *RedisStorage) Store(url string) (key string, err error) {
	return
}
func (s *RedisStorage) Load(key string) (url string, err error) {
	return
}
