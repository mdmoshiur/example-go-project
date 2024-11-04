package cache

import (
	"context"
	"strings"
	"time"

	"github.com/mdmoshiur/example-go/internal/conn"
	"github.com/redis/go-redis/v9"
)

// RedisCache represents the redis driver
type RedisCache struct {
	rds        conn.RedisClient
	prefix     string
	defaultTTL time.Duration
}

var ctx = context.Background()

// NewRedisCache is the factory function for the redis cache instance
func NewRedisCache(r conn.RedisClient, pfx string, dtl time.Duration) Cacher {
	return &RedisCache{
		rds:        r,
		prefix:     pfx,
		defaultTTL: dtl,
	}
}

// SetCache set cache on redis
func (r *RedisCache) SetCache(key string, val string, ttl time.Duration) error {
	err := r.rds.Set(ctx, r.prefix+key, val, ttl).Err()

	return err
}

// GetCache returns a cache value against the key provided from redis
func (r *RedisCache) GetCache(key string) (string, error) {
	val, err := r.rds.Get(ctx, r.prefix+key).Result()
	if err != nil {
		if err == redis.Nil { // key does not exist
			return "", nil
		}

		return "", err
	}

	return val, nil
}

// ClearCache clears a cache matching the pattern of the redis key
func (r *RedisCache) ClearCache(pattern string) error {
	pattern = r.prefix + pattern
	keys, err := r.rds.Keys(ctx, pattern).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}

		return err
	}

	if len(keys) > 0 {
		if err = r.rds.Del(ctx, keys...).Err(); err != nil {
			return err
		}
	}

	return nil
}

// Exists checks if a redis key exists or not
func (r *RedisCache) Exists(key string) (bool, error) {
	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	_, err := r.rds.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// BuildKey builds a valid key for the redis cache
func (r *RedisCache) BuildKey(keywords ...string) string {
	pfx := []string{
		strings.TrimSuffix(r.prefix, "_"),
	}

	for _, kw := range keywords {
		if kw != "" && kw != r.prefix {
			pfx = append(pfx, kw)
		}
	}
	key := strings.Join(pfx, "_")
	return key
}
