package main

import (
	ConfigPkg "github.com/dragoscojocaru/forxy/internal/config"
	HttpServer "github.com/dragoscojocaru/forxy/internal/server"
)

func main() {
	server := HttpServer.NewServer()
	server.Serve(ConfigPkg.Configuration.Server.Bind_Port)
}
