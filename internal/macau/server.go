package macau

import (
	"github.com/go-redis/redis"
)

type Server struct {
	redis *redis.Client
}

func (s *Server) Run(addr string) error {
	// handle websockets
	s.Listen(addr)

	return nil
}
