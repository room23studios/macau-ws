package macau

import (
	"github.com/go-redis/redis"
)

type Server struct {
	redis *redis.Client
}

func (s *Server) Run(addr string) error {
	// connect to redis
	redisOptions := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	redisClient, err := SetupRedis(redisOptions)
	if err != nil {
		return err
	}
	s.redis = redisClient

	// handle websockets
	s.Listen(addr)

	return nil
}
