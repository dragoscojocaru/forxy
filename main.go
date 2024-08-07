package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {

	http.HandleFunc("/http/sequential", SequentialHTTPHandler)
	http.HandleFunc("/http/fork", ForkHTTPHandler)
	http.HandleFunc("/http/test", TestHandler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func SequentialHTTPHandler(w http.ResponseWriter, _ *http.Request) {

	client := &http.Client{}

	start := time.Now()

	for range 3 {
		req, err1 := http.NewRequest("GET", "https://google.com/", nil)
		resp, err2 := client.Do(req)
		if err1 == nil && err2 == nil {
			fmt.Fprintf(w, resp.Status)
		} else {
			fmt.Fprintf(w, err2.Error())
		}
	}

	fmt.Println("Execution time on sequential: ", time.Since(start))

}

func ForkHTTPHandler(w http.ResponseWriter, _ *http.Request) {

	client := &http.Client{}

	var wg sync.WaitGroup

	start := time.Now()

	for range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, err1 := http.NewRequest("GET", "https://google.com/", nil)
			resp, err2 := client.Do(req)
			if err1 == nil && err2 == nil {
				fmt.Fprintf(w, resp.Status)
			} else {
				fmt.Fprintf(w, err2.Error())
			}
		}()
	}

	wg.Wait()

	fmt.Println("Execution time on fork: ", time.Since(start))
}

type RequestMessage struct {
	URL    string
	Method string
	Body   json.RawMessage
}

type ResponseMessage struct {
	Body []json.RawMessage
}

type ForxyBodyPayload struct {
	Timeout  int
	Requests []RequestMessage
}

func HTTPRequest(requestMessage RequestMessage, client *http.Client, ch *chan http.Response, wg *sync.WaitGroup) {
	defer wg.Done()
	//bodyReader := bytes.NewReader([]byte(requestMessage.Body))

	req, err1 := http.NewRequest(requestMessage.Method, requestMessage.URL, nil)
	//req, err1 := http.NewRequest("get", "https://google.com/", nil)
	resp, err2 := client.Do(req)
	if err1 != nil && err2 != nil {
		//TODO
		fmt.Println("error handling")
		os.Exit(1)
	}

	*ch <- *resp
}

func TestHandler(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}

	decoder := json.NewDecoder(r.Body)

	var t ForxyBodyPayload
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	responseChannel := make(chan http.Response, len(t.Requests))

	var wg sync.WaitGroup

	for idx := range t.Requests {
		wg.Add(1)
		go HTTPRequest(t.Requests[idx], client, &responseChannel, &wg)
	}

	wg.Wait()
	fmt.Println(1)

	for idx := 0; idx < len(t.Requests); idx++ {
		rs := <-responseChannel
		fmt.Fprintf(w, rs.Status)
	}
}
