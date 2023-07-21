package server

import (
	"net/http"

	"github.com/wb/cmd/0L/internal/server/public"
)

func router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/json", public.Distribute)

	return mux
}
