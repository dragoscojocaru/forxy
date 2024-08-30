package http

import (
	ForxyHttp "github.com/dragoscojocaru/forxy/handler/http"
	"log"
	"net/http"
	"strconv"
)

type Server struct{}

func (server *Server) Serve(bindPort int) {

	http.HandleFunc("/http/sequential", ForxyHttp.HTTPSequentialHandler)
	http.HandleFunc("/http/fork", ForxyHttp.ForkHandler)

	log.Println("Forxy HTTP server listening for connections on port " + strconv.Itoa(bindPort) + " ...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(bindPort), nil))
}
