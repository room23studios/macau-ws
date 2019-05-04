package main

import (
	"github.com/Room23Studios/macau-ws/internal/macau"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := macau.Server{}
	panic(server.Run(":1234"))
}
