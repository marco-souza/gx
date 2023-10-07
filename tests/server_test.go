package main

import (
	"testing"

	"github.com/marco-souza/gx/cmd/server"
)

var (
	hostname = "localhost"
	port     = 3001
)

func TestCanCreateServer(t *testing.T) {
	s := server.New()
	if s == nil {
		t.Fatal("Failed to create server")
	}
}
