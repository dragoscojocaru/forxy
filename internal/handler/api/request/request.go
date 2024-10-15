package request

import (
	"encoding/json"
	"time"
)

type RequestMessage struct {
	URL     string
	Method  string
	Body    json.RawMessage
	Headers map[string]string
}

type ForxyBodyPayload struct {
	Timeout  time.Duration
	Requests map[int]RequestMessage
}
