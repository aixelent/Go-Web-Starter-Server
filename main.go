package main

import (
	"./system/app"
	"flag"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "Assigning server port for listening:")
	flag.Parse()
}

func main() {
	server := app.NewServer()
	server.Init(port)
	server.Start()
}
