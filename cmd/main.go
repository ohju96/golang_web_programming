package main

import (
	"golang_web_programming/internal/membership"
)

func main() {
	server := membership.NewDefaultServer()
	server.Run()
}
