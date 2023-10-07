package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/marco-souza/gx/cmd/server"
)

func TestCanCreateServer(t *testing.T) {
	s := server.New()
	values := fmt.Sprintf("%v", s)

	t.Run("validate env configs", func(t *testing.T) {
		if s == nil {
			t.Fatal("Failed to create server")
		}
		if !strings.Contains(values, "localhost") {
			t.Fatal("No host found")
		}
		if !strings.Contains(values, "3001") {
			t.Fatal("No port found")
		}
	})
}
