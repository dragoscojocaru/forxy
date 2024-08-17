package response

import (
	"encoding/json"
	"fmt"
	"github.com/dragoscojocaru/forxy/logger"
	"io"
	"log"
	"net/http"
)

type ForxyResponsePayload struct {
	Responses map[int]json.RawMessage `json:"responses"`
}

type ForxyResponsePayloadWriter struct {
	http.ResponseWriter
}

func NewForxyResponsePayload() *ForxyResponsePayload {
	responses := make(map[int]json.RawMessage)
	forxyResponsePayload := ForxyResponsePayload{
		Responses: responses,
	}

	fmt.Println(forxyResponsePayload)

	return &forxyResponsePayload
}

func (forxyResponsePayload *ForxyResponsePayload) AddResponse(idx int, response http.Response) {

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	forxyResponsePayload.Responses[idx] = bodyBytes
}

func NewForxyPayloadWriter() *ForxyResponsePayloadWriter {
	return new(ForxyResponsePayloadWriter)
}

func (ForxyResponsePayloadWriter) JsonMarshal(w http.ResponseWriter, payload ForxyResponsePayload) {

	bytes, err := json.Marshal(payload)
	if err != nil {
		logger.FileErrorLog(err)
	}
	in, err := w.Write(bytes)
	fmt.Println(in)
	if err != nil {
		logger.FileErrorLog(err)
	}

}

type ResponseInternalChannel struct {
	idx      int
	response http.Response
}

func NewResponseInternalChannel(idx int, response http.Response) *ResponseInternalChannel {
	tmp := new(ResponseInternalChannel)

	tmp.idx = idx
	tmp.response = response

	return tmp
}

func GetIdx(responseInternalChannel *ResponseInternalChannel) int {
	return responseInternalChannel.idx
}

func GetResponse(responseInternalChannel *ResponseInternalChannel) http.Response {
	return responseInternalChannel.response
}
