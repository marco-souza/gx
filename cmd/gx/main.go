package main

import "github.com/marco-souza/gx/cmd/server"

func main() {
	s := server.New("localhost", 8001)
	s.Start()
}
