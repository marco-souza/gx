package main

import "github.com/marco-souza/gx/cmd/server"

func main() {
	s := server.New("localhost", 3001)
	s.Start()
}
