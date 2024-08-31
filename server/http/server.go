package http

import (
	"fmt"
	ForxyHttp "github.com/dragoscojocaru/forxy/handler/http"
	"log"
	"net/http"
)

type Server struct{}

func (server *Server) Serve(bindPort int32) {

	http.HandleFunc("/http/sequential", ForxyHttp.HTTPSequentialHandler)
	http.HandleFunc("/http/fork", ForxyHttp.ForkHandler)

	port := fmt.Sprint(bindPort)

	log.Println("Forxy HTTP server listening for connections on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
