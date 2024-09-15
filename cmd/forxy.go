package main

import (
	ConfigPkg "github.com/dragoscojocaru/forxy/pkg/config"
	HttpServer "github.com/dragoscojocaru/forxy/pkg/server/http"
)

func main() {

	server := HttpServer.Server{}
	server.Serve(ConfigPkg.Configuration.Server.Bind_Port)
}
