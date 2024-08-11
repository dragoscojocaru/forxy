package http

import (
	forxy_http "github.com/dragoscojocaru/forxy/handler/http"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/http/sequential", forxy_http.HTTPSequentialHandler)
	http.HandleFunc("/http/fork", forxy_http.HTTPForkHandler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
