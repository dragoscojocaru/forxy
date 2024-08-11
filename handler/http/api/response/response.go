package response

import (
	"io"
	"net/http"
	"strconv"
	"strings"
)

type ResponseMessage struct {
	Body map[int]io.ReadCloser
}

func NewResponseMessage() *ResponseMessage {

	body := make(map[int]io.ReadCloser)
	rm := new(ResponseMessage)
	rm.Body = body

	return rm
}

func AddResponse(responseMessage *ResponseMessage, idx int, response *http.Response) {
	responseMessage.Body[idx] = response.Body
}

func GetResponseStream(message *ResponseMessage) *[]io.Reader {
	var stream []io.Reader
	for idx := range message.Body {
		stream = append(stream, strings.NewReader(strconv.Itoa(idx)+": "))
		stream = append(stream, message.Body[idx])
	}

	return &stream
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

func GetResponse(responseInternalChannel *ResponseInternalChannel) *http.Response {
	return &responseInternalChannel.response
}
