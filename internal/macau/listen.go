package macau

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Room23Studios/macau-ws/internal/proto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Server) Listen(addr string) {
	http.HandleFunc("/socket", s.Handler)
	log.Printf("Listening at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type HelloClaims struct {
	GameID string `json:"game_id"`
	Nick   string `json:"nick"`
	jwt.StandardClaims
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
		case *proto.CommandHello:
			t, err := jwt.ParseWithClaims(c.Token, &HelloClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				log.Println(err)
				return // TODO: send an error
			}

			claims := t.Claims.(*HelloClaims)
			log.Printf("Game %s: '%s' joins.\n", claims.GameID, claims.Nick)

			gameInterface, loaded := s.games.LoadOrStore(claims.GameID, &Game{
				ID:      claims.GameID,
				Players: []string{},
			})

			game := gameInterface.(*Game)
			conn.WriteJSON(map[string]interface{}{"command": "hello", "data": map[string]interface{}{}})
			// TODO: start a game goroutine and pass events there

			log.Printf("game %s, loaded? %t, nicks: %s", game.ID, loaded, game.Players)
		}
	}
}
