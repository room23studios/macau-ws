package macau

import (
	"sync"

	"github.com/go-redis/redis"
)

type Server struct {
	redis *redis.Client
	games sync.Map
}

func (s *Server) Run(addr string) error {
	// handle websockets
	s.Listen(addr)

	return nil
}
