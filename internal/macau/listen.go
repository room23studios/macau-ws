package macau

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func (s *Server) Listen(addr string) {
	http.HandleFunc("/socket", s.Handler)
	fmt.Printf("Listening at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
