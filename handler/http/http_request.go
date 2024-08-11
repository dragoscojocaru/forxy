package http

import (
	"bytes"
	forxy_http_api_request "github.com/dragoscojocaru/forxy/handler/http/api/request"
	"net/http"
	"os"
	"sync"
)

func HTTPRequest(requestMessage forxy_http_api_request.RequestMessage, client *http.Client, ch *chan http.Response, wg *sync.WaitGroup) {
	defer wg.Done()

	bodyReader := bytes.NewReader([]byte(requestMessage.Body))

	req, err1 := http.NewRequest(requestMessage.Method, requestMessage.URL, bodyReader)

	for key, value := range requestMessage.Headers {
		req.Header.Set(key, value)
	}

	resp, err2 := client.Do(req)

	if err1 != nil && err2 != nil {
		//TODO implement error handling
		os.Exit(1)
	}

	*ch <- *resp
}
