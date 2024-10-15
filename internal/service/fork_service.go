package service

import (
	"encoding/json"
	"fmt"
	ForxyHttpApiRequest "github.com/dragoscojocaru/forxy/internal/handler/api/request"
	"github.com/dragoscojocaru/forxy/internal/handler/api/response"
	forxyHttp "github.com/dragoscojocaru/forxy/internal/handler/request"
	"github.com/dragoscojocaru/forxy/internal/logger"
	"io"
	"net/http"
	"time"
)

type ForkService interface {
	Fork(w http.ResponseWriter, r *http.Request)
}

type forkService struct {
	payloadWriter *response.ForxyResponsePayloadWriter
}

func NewForkService() ForkService {
	return &forkService{
		payloadWriter: response.NewForxyPayloadWriter(),
	}
}

func (fs *forkService) Fork(w http.ResponseWriter, r *http.Request) {
	body := fs.getRequestBody(r)

	responseChannel := make(chan response.ChannelMessage, len(body.Requests))
	ech := make(chan bool, 1)

	go forxyHttp.SendStream(&responseChannel, body, &ech)

	responsePayload := new(response.ForxyResponsePayload)

	select {
	case <-time.After(body.Timeout * time.Millisecond):
		responsePayload.Message = fmt.Sprintf("Request timed out after %d miliseconds.", body.Timeout)
	case _ = <-ech:
		responsePayload = fs.computePayload(body, responseChannel)
	}

	fs.payloadWriter.JsonMarshal(w, *responsePayload)
}

func (fs *forkService) computePayload(body ForxyHttpApiRequest.ForxyBodyPayload,
	responseChannel chan response.ChannelMessage,
) *response.ForxyResponsePayload {
	responsePayload := response.NewForxyResponsePayload()

	for _ = range body.Requests {
		rs := <-responseChannel
		res := response.GetResponse(&rs)

		responsePayload.AddResponse(response.GetIdx(&rs), res)
	}

	return responsePayload
}

func (fs *forkService) getRequestBody(r *http.Request) ForxyHttpApiRequest.ForxyBodyPayload {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.FileErrorLog(err)
	}

	var body ForxyHttpApiRequest.ForxyBodyPayload
	err = json.Unmarshal(bodyBytes, &body)

	return body
}
