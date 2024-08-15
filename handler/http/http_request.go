package http

import (
	"bytes"
	ForxyHttpApiRequest "github.com/dragoscojocaru/forxy/handler/http/api/request"
	"github.com/dragoscojocaru/forxy/handler/http/api/response"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

func HTTPRequest(idx int, requestMessage ForxyHttpApiRequest.RequestMessage, ch *chan response.ResponseInternalChannel, wg *sync.WaitGroup) {
	defer wg.Done()

	bodyReader := bytes.NewReader(requestMessage.Body)

	req, err1 := http.NewRequest(requestMessage.Method, requestMessage.URL, bodyReader)

	for key, value := range requestMessage.Headers {
		req.Header.Set(key, value)
	}

	client := connectionPool.GetServerConnection(GetHost(requestMessage.URL))
	resp, err2 := client.Do(req)

	if err1 != nil && err2 != nil {
		//TODO implement error handling
		os.Exit(1)
	}

	chanResp := response.NewResponseInternalChannel(idx, *resp)

	*ch <- *chanResp
}

func SendStream(ch *chan response.ResponseInternalChannel, body ForxyHttpApiRequest.ForxyBodyPayload) {
	var wg sync.WaitGroup

	for idx := range body.Requests {
		wg.Add(1)
		go HTTPRequest(idx, body.Requests[idx], ch, &wg)
	}

	wg.Wait()
}

func GetHost(link string) string {
	urlS, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	return urlS.Hostname()
}
