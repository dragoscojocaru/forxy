package main

import (
	ConfigPkg "github.com/dragoscojocaru/forxy/config"
	HttpServer "github.com/dragoscojocaru/forxy/server/http"
)

func main() {

	//TODO implement file error handling trough the project

	server := HttpServer.Server{}
	server.Serve(ConfigPkg.Configuration.Server.Bind_Port)

}
