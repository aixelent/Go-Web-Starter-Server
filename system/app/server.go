package app

import (
	"log"
	"net/http"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

// Init values
func (s *Server) Init(port string) {
	log.Println("Initializing server...")
	s.port = ":" + port
}

// Start server
func (s *Server) Start() {
	log.Println("Server starting on port:", s.port)
	http.ListenAndServe(s.port, nil)
}
