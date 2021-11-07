package server

import (
	"dev-1/controller"
	"net/http"
)

type Server struct {
	S *http.ServeMux
}

func (s *Server) Initialize() {
	s.S = http.NewServeMux()
	s.S.HandleFunc("/hello", controller.Hello)
	s.S.HandleFunc("/readfile", controller.ReadFile)
	s.S.HandleFunc("/writetofile", controller.WriteToFile)
	_ = http.ListenAndServe(":8080", s.S)
}
