package server

import (
	"context"
	"net/http"
	"time"

	"github.com/wb/cmd/0L/internal/config"
)

const (
	ErrRun = "server run error"
)

type Server struct {
	srv *http.Server
}

var Srv *Server

func NewServer() *Server {
	return &Server{
		srv: &http.Server{
			Addr:    config.App.Server.Host + ":" + config.App.Server.Port,
			Handler: router(),
		},
	}
}

func (s *Server) Run() (err error) {
	// defer func() { err = erro.IsError(ErrRun, err) }()
	err = s.srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Close() {
	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	s.srv.Shutdown(ctxShutdown)
}

func init() {
	Srv = NewServer()
}
