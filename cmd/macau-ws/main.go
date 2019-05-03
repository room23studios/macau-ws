package main

import (
	"fmt"

	"github.com/Room23Studios/macau-ws/internal/macau"
)

func main() {
	fmt.Println("Goodbye, world!")

	server := macau.Server{}
	panic(server.Run(":1234"))
}
