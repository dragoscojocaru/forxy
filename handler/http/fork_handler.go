package http

import (
	"encoding/json"
	ForxyHttpApiRequest "github.com/dragoscojocaru/forxy/handler/http/api/request"
	"github.com/dragoscojocaru/forxy/handler/http/api/response"
	"net/http"
)

func ForkHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var body ForxyHttpApiRequest.ForxyBodyPayload
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	responseChannel := make(chan response.ResponseInternalChannel, len(body.Requests))

	SendStream(&responseChannel, body)

	forxyResponsePayload := response.NewForxyResponsePayload()
	for idx := range body.Requests {

		rs := <-responseChannel
		res := response.GetResponse(&rs)

		forxyResponsePayload.AddResponse(idx, res)

	}
	forxyPayloadWriter := response.NewForxyPayloadWriter()
	forxyPayloadWriter.JsonMarshal(w, *forxyResponsePayload)
}
