package main

import (
	"github.com/wb/cmd/0L/internal/server"
)

func main() {
	if err := server.NewServer().Run(); err != nil {
		panic(err)
	}
}
