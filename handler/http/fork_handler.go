package http

import (
	"encoding/json"
	ForxyHttpApiRequest "github.com/dragoscojocaru/forxy/handler/http/api/request"
	"github.com/dragoscojocaru/forxy/handler/http/api/response"
	"github.com/dragoscojocaru/forxy/logger"
	"net/http"
)

func ForkHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var body ForxyHttpApiRequest.ForxyBodyPayload
	err := decoder.Decode(&body)
	if err != nil {
		go logger.FileErrorLog(err)
	}

	responseChannel := make(chan response.ChannelMessage, len(body.Requests))

	SendStream(&responseChannel, body)

	forxyResponsePayload := response.NewForxyResponsePayload()
	for _ = range body.Requests {

		rs := <-responseChannel
		res := response.GetResponse(&rs)

		forxyResponsePayload.AddResponse(response.GetIdx(&rs), res)

	}
	forxyPayloadWriter := response.NewForxyPayloadWriter()
	forxyPayloadWriter.JsonMarshal(w, *forxyResponsePayload)
}
