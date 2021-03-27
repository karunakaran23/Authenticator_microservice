package server

import (
	"authentication_microservice/handler"
	"authentication_microservice/middleware"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

func (s *Server) InitRoute(h *handler.Handler) {
	s.Router.HandleFunc("/FindUser", middleware.JSONandCORS(h.FindUser))
	s.Router.HandleFunc("/auth", middleware.JSONandCORS(h.Auth))

}
