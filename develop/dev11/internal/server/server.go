package server

import (
	"log"
	"net/http"
	"wb_l2/develop/dev11/internal/config"
	"wb_l2/develop/dev11/internal/route"
)

type Server struct {
	cfg *config.Config
	mux *route.Mux
}

func NewServer(mux *route.Mux, cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
		mux: mux,
	}
}

func (s *Server) Run() {

	rt := s.mux.SetRoutes()

	if err := http.ListenAndServe(s.cfg.Port, rt); err != nil {
		log.Fatalln(err)
	}
}
