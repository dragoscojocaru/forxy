package main

import (
	ForxyHttp "github.com/dragoscojocaru/forxy/handler/http"
	ForxyLogger "github.com/dragoscojocaru/forxy/logger"
	"log"
	"net/http"
	"os"
)

func main() {

	//TODO implement file error handling trough the project
	_, err := os.OpenFile("forceError", os.O_RDONLY, 111)
	if err != nil {
		ForxyLogger.FileErrorLog(err)
	}

	http.HandleFunc("/http/sequential", ForxyHttp.HTTPSequentialHandler)
	http.HandleFunc("/http/fork", ForxyHttp.ForkHandler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
