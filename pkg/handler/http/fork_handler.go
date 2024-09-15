package http

import (
	"encoding/json"
	ForxyHttpApiRequest "github.com/dragoscojocaru/forxy/pkg/handler/http/api/request"
	"github.com/dragoscojocaru/forxy/pkg/handler/http/api/response"
	"github.com/dragoscojocaru/forxy/pkg/logger"
	"io/ioutil"
	"net/http"
)

func ForkHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.FileErrorLog(err)
	}

	var body ForxyHttpApiRequest.ForxyBodyPayload
	err = json.Unmarshal(bodyBytes, &body)

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
