package http

import "net/http"

type ClientConnectionPool struct {
	ServerHashmap map[string]*http.Client
}

func NewClientConnectionPool() *ClientConnectionPool {
	serverHashmap := make(map[string]*http.Client)

	clientConnectionPool := ClientConnectionPool{
		serverHashmap,
	}

	return &clientConnectionPool
}

func (pool ClientConnectionPool) GetServerConnection(serverHost string) *http.Client {
	if pool.ServerHashmap[serverHost] == nil {
		pool.ServerHashmap[serverHost] = &http.Client{}
	}
	return pool.ServerHashmap[serverHost]
}

var connectionPool = NewClientConnectionPool()
