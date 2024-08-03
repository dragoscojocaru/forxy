package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {

	http.HandleFunc("/http/sequential", SequentialHTTPHandler)
	http.HandleFunc("/http/fork", ForkHTTPHandler)

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
