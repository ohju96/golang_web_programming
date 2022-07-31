package main

import (
	"golang_web_programming/internal"
)

// @title Membership API
// @version 1.0
// @host localhost:8080
func main() {
	server := internal.NewDefaultServer()
	server.Run()
}
