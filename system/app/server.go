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
func (serv *Server) Init(port string) {
	log.Println("Initializing server...")
	serv.port = ":" + port
}

// Start server
func (serv *Server) Start() {
	log.Println("Server starting on port:", serv.port)
	log.Println("http://localhost" + serv.port)
	http.ListenAndServe(serv.port, nil)
}
