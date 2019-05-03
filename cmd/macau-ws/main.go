package main

import (
	"github.com/Room23Studios/macau-ws/internal/macau"
)

func main() {
	server := macau.Server{}
	panic(server.Run(":1234"))
}
