package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

// Config redis.
type Config struct {
	Host string
	Port int
}

// LoadEnv - load configuration from env.
func LoadEnv() redis.Options {
	return redis.Options{
		Addr:     fmt.Sprintf("%s:%d", os.Getenv("REDIS_HOST"), port()),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db(),
	}
}

// NewRedis creates new connection to redis and return the connection
func NewRedis(cfg redis.Options) (*redis.Client, context.Context, error) {
	conn := redis.NewClient(&cfg)
	ctx := context.Background()
	_, err := conn.Ping(ctx).Result()
	if err != nil {
		return nil, nil, err
	}

	// Setup redis to send keyspace events
	_, err = conn.ConfigSet(ctx, "notify-keyspace-events", "KEt").Result()
	if err != nil {
		return nil, nil, err
	}

	return conn, ctx, nil
}

func port() int {
	p, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return 6379
	}
	return p
}

func db() int {
	d, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return 0
	}
	return d
}
