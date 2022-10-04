package main

import (
	"fmt"

	"github.com/polpettone/go-samples/http-server/server"
)

func main() {
	fmt.Println("start server")
	RunServer()

	server.RunServer()
}
