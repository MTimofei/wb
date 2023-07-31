package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/wb/cmd/0L/internal/server"
	"github.com/wb/cmd/0L/internal/serves"
	"github.com/wb/pkg/logg"
)

var wg sync.WaitGroup

func main() {
	if err := serves.App.Start(); err != nil {
		serves.App.Log.Error("start serves", logg.Err(err))
		os.Exit(1)
	}
	wg.Add(3)

	go func() {
		defer wg.Done()
		serves.App.Work()
	}()

	go func() {
		defer wg.Done()
		gracefulShutdown()
	}()

	go func() {
		defer wg.Done()
		if err := server.Srv.Run(); err != nil && err != http.ErrServerClosed {
			serves.App.Log.Error("server", logg.Err(err))
		}
	}()

	wg.Wait()
	fmt.Println("app stope")
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	server.Srv.Close()
	serves.App.Close()
}
