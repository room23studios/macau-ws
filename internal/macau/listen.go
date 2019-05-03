package macau

import (
	"log"
	"net/http"

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
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(messageType, p, err)
	}
}
