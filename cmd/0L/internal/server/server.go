package server

import (
	"net/http"

	"github.com/wb/cmd/0L/internal/config"
	"github.com/wb/pkg/erro"
)

const (
	ErrRun = "server run error"
)

type Server struct {
	srv *http.Server
}

func NewServer() *Server {
	return &Server{
		srv: &http.Server{
			Addr:    config.App.Server.Host + ":" + config.App.Server.Port,
			Handler: router(),
		},
	}
}

func (s *Server) Run() error {
	defer erro.IsError(ErrRun, s.srv.ListenAndServe())
	err := s.srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
