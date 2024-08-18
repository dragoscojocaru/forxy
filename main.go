package main

import (
	HttpServer "github.com/dragoscojocaru/forxy/server/http"
)

func main() {

	//TODO implement file error handling trough the project
	//_, err := os.OpenFile("forceError", os.O_RDONLY, 111)
	//if err != nil {
	//	ForxyLogger.FileErrorLog(err)
	//}

	server := HttpServer.Server{}
	server.Serve(8080)

}
