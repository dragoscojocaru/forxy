package request

import "encoding/json"

type RequestMessage struct {
	URL     string
	Method  string
	Body    json.RawMessage
	Headers map[string]string
}

type ResponseMessage struct {
	Body []json.RawMessage
}

type ForxyBodyPayload struct {
	Timeout  int
	Requests []RequestMessage
}
