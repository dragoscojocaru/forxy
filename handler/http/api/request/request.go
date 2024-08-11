package request

import "encoding/json"

type RequestMessage struct {
	URL     string
	Method  string
	Body    json.RawMessage
	Headers map[string]string
}

type ForxyBodyPayload struct {
	Timeout  int
	Requests map[int]RequestMessage
}
