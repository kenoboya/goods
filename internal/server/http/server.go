package http

import "net/http"

type Server struct {
	srv  *http.Server
}

func NewServer(config)