package cache

import "time"

type Cacher interface {
	SetCache(key string, val string, ttl time.Duration) error
	GetCache(key string) (string, error)
	ClearCache(pattern string) error
	Exists(key string) (bool, error)
	BuildKey(keywords ...string) string
}
