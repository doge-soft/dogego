package cache

import (
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

var CacheClient *redis.Client

func ConnectRedisCache() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)

	CacheClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       int(db),
	})
}
