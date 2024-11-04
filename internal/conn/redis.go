package conn

import (
	"context"
	"github.com/mdmoshiur/example-go/internal/config"
	"github.com/redis/go-redis/v9"
)

// RedisClient holds the redis client instance
type RedisClient struct {
	*redis.Client
}

// RedisCl is an instance of *redisClient{}
var redisCl RedisClient

// Connect assigns redis.Client interface base on config to RedisClient
func (r *RedisClient) Connect(cfg *config.RedisCfg) error {
	rc := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := rc.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	r.Client = rc

	return nil
}

// ConnectRedis provides a connector to redis based on configurations set
func ConnectRedis() error {
	cfg := config.Redis()
	err := redisCl.Connect(&cfg)

	return err
}

// DefaultRedis returns the default RedisClient currently in use
func DefaultRedis() RedisClient {
	return redisCl
}
