package macau

import (
	"fmt"

	"github.com/go-redis/redis"
)

func SetupRedis(opt *redis.Options) (*redis.Client, error) {
	r := redis.NewClient(opt)

	_, err := r.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("redis connection failed")
	}

	return r, nil
}
