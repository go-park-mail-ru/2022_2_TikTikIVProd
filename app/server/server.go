package server

import (
	"net/http"
	"time"
	"log"
)

type Server struct {
	http.Server
	/*тут конфиги и логгер ещё будет*/
}

func NewServer(r http.Handler) *Server {
	return &Server{
		http.Server{
			Addr:              ":8080",
			Handler:           r,
			ReadTimeout:       10 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			WriteTimeout:      10 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	log.Println("start serving :8080")
	return s.ListenAndServe()
}

