package http

import (
	forxy_http "github.com/dragoscojocaru/forxy/handler/http"
	"log"
	"net/http"
)

type Server struct{}

func (server *Server) Serve(bindPort int) {
	http.HandleFunc("/http/sequential", forxy_http.HTTPSequentialHandler)
	http.HandleFunc("/http/fork", forxy_http.ForkHandler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
