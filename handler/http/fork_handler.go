package http

import (
	"encoding/json"
	"fmt"
	forxy_http_api_request "github.com/dragoscojocaru/forxy/handler/http/api/request"
	"net/http"
	"sync"
)

func HTTPForkHandler(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}

	decoder := json.NewDecoder(r.Body)

	var body forxy_http_api_request.ForxyBodyPayload
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	responseChannel := make(chan http.Response, len(body.Requests))

	var wg sync.WaitGroup

	for idx := range body.Requests {
		wg.Add(1)
		go HTTPRequest(body.Requests[idx], client, &responseChannel, &wg)
	}

	wg.Wait()

	for idx := 0; idx < len(body.Requests); idx++ {
		rs := <-responseChannel
		fmt.Fprintf(w, rs.Status)
	}
}
