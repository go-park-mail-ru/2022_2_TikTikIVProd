package server

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/router"
	"log"
	"net/http"
	"time"
)

type Server struct {
	http.Server
	/*тут конфиги и логгер ещё будет*/
}

func NewServer(e *router.EchoRouter) *Server {
	return &Server{
		http.Server{
			Addr:              ":8080",
			Handler:           e,
			ReadTimeout:       10 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			WriteTimeout:      10 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	log.Println("start serving in :8080")
	return s.ListenAndServe()
}
