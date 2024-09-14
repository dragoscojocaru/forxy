package http

import (
	"bytes"
	"encoding/json"
	ForxyHttpApiRequest "github.com/dragoscojocaru/forxy/handler/http/api/request"
	"github.com/dragoscojocaru/forxy/handler/http/api/response"
	"github.com/dragoscojocaru/forxy/logger"
	"net/http"
	"net/url"
	"sync"
)

func HTTPRequest(idx int, requestMessage ForxyHttpApiRequest.RequestMessage, ch *chan response.ChannelMessage, wg *sync.WaitGroup) {
	defer wg.Done()

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(requestMessage.Body)

	req, err1 := http.NewRequest(requestMessage.Method, requestMessage.URL, &buf)
	if err1 != nil {
		logger.FileErrorLog(err1)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range requestMessage.Headers {
		req.Header.Set(key, value)
	}

	host, err := GetHost(requestMessage.URL)
	if err != nil {
		logger.FileErrorLog(err)
	}

	client := connectionPool.GetServerConnection(host)
	resp, err2 := client.Do(req)

	if err1 != nil && err2 != nil {
		go logger.FileErrorLog(err1)
		go logger.FileErrorLog(err2)
	}

	chanResp := response.NewChannelMessage(idx, *resp)

	*ch <- *chanResp
}

func SendStream(ch *chan response.ChannelMessage, body ForxyHttpApiRequest.ForxyBodyPayload) {
	var wg sync.WaitGroup

	for idx := range body.Requests {
		wg.Add(1)
		go HTTPRequest(idx, body.Requests[idx], ch, &wg)
	}
	wg.Wait()
}

func GetHost(link string) (string, error) {
	urlS, err := url.Parse(link)
	if err != nil {
		go logger.FileErrorLog(err)
		return "", err
	}
	return urlS.Hostname(), nil
}
