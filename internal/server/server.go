package server

import (
	"fmt"
	"github.com/dragoscojocaru/forxy/internal/handler"
	"log"
	"net/http"
)

type Server interface {
	Serve(bindPort int32)
}

type server struct {
	forkHandler handler.ForkHandler
}

func NewServer() Server {
	return &server{
		forkHandler: handler.NewForkHandler(),
	}
}

func (server *server) Serve(bindPort int32) {
	http.HandleFunc("/http/sequential", handler.HTTPSequentialHandler)
	http.HandleFunc("/http/fork", server.forkHandler.Handle)

	port := fmt.Sprint(bindPort)

	log.Println("Forxy HTTP server listening for connections on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
