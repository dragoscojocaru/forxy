package http

import (
	"encoding/json"
	ForxyHttpApiRequest "github.com/dragoscojocaru/forxy/handler/http/api/request"
	"github.com/dragoscojocaru/forxy/handler/http/api/response"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ForkHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var body ForxyHttpApiRequest.ForxyBodyPayload
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	responseChannel := make(chan response.ResponseInternalChannel, len(body.Requests))

	SendStream(&responseChannel, body)

	responseMessage := response.NewResponseMessage()

	//TODO implement response stream structure
	_, err = io.Copy(w, strings.NewReader("{\"responses\": {"))
	var i = 0
	for idx := range body.Requests {

		rs := <-responseChannel

		indexReader := strings.NewReader(commaIndex(i) + "\"" + strconv.Itoa(idx) + "\": ")
		combinedReader := io.MultiReader(indexReader, response.GetResponse(&rs).Body)

		_, err = io.Copy(w, combinedReader)

		response.AddResponse(responseMessage, idx, response.GetResponse(&rs))
		i++
	}
	_, err = io.Copy(w, strings.NewReader("}}"))
}

func commaIndex(idx int) string {
	if idx > 0 {
		return ","
	}
	return ""
}
