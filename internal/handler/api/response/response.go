package response

import (
	"encoding/json"
	"github.com/dragoscojocaru/forxy/internal/handler/api"
	"github.com/dragoscojocaru/forxy/internal/logger"
	"io"
	"net/http"
)

type ServerResponse struct {
	Control api.Control     `json:"forxy_control"`
	Status  int             `json:"status"`
	Body    json.RawMessage `json:"body"`
}

func NewServerResponse(response http.Response) *ServerResponse {
	control := api.NewControl()

	serverResponse := ServerResponse{
		Status: 500,
		Body:   []byte("{}"),
	}
	serverResponse.Control = *control
	serverResponse.Control.Validate(response)
	serverResponse.Status = response.StatusCode

	if serverResponse.Control.Ok == true {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			go logger.FileErrorLog(err)
		}
		serverResponse.Body = bodyBytes
	}

	return &serverResponse
}

type ForxyResponsePayload struct {
	Responses map[int]*ServerResponse `json:"responses"`
}

type ForxyResponsePayloadWriter struct {
	http.ResponseWriter
}

func NewForxyResponsePayload() *ForxyResponsePayload {
	responses := make(map[int]*ServerResponse)
	forxyResponsePayload := ForxyResponsePayload{
		Responses: responses,
	}

	return &forxyResponsePayload
}

func (forxyResponsePayload *ForxyResponsePayload) AddResponse(idx int, response http.Response) {
	forxyResponsePayload.Responses[idx] = NewServerResponse(response)
}

func NewForxyPayloadWriter() *ForxyResponsePayloadWriter {
	return new(ForxyResponsePayloadWriter)
}

func (ForxyResponsePayloadWriter) JsonMarshal(w http.ResponseWriter, payload ForxyResponsePayload) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		go logger.FileErrorLog(err)
	}
	_, err = w.Write(bytes)
	if err != nil {
		go logger.FileErrorLog(err)
	}
}
