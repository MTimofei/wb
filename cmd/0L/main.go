package main

import (
	"github.com/wb/cmd/0L/internal/server"
	"github.com/wb/cmd/0L/internal/serves"
)

func main() {
	if err := serves.App.Start(); err != nil {
		panic(err)
	}

	go serves.App.Work()

	if err := server.NewServer().Run(); err != nil {
		panic(err)
	}
}
