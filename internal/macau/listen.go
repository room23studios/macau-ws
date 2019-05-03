package macau

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Room23Studios/macau-ws/internal/proto"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *Server) Listen(addr string) {
	http.HandleFunc("/socket", s.Handler)
	log.Printf("Listening at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		command, err := proto.ParseMessage(p)
		if err != nil {
			log.Println(err)
			return
		}

		switch c := command.(type) {
		case *proto.CommandPing:
			fmt.Println("got pinged:", c.Payload)
		case *proto.CommandJoin:
			fmt.Println("join received:", c.GamePIN)
		}
	}
}
