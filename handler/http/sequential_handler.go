package http

import (
	"fmt"
	"net/http"
	"time"
)

func HTTPSequentialHandler(w http.ResponseWriter, _ *http.Request) {

	client := &http.Client{}

	start := time.Now()

	for range 5 {
		req, err1 := http.NewRequest("GET", "https://catalog.dedeman.ro/api/live/list", nil)
		resp, err2 := client.Do(req)
		if err1 == nil && err2 == nil {
			fmt.Fprintf(w, resp.Status)
		} else {
			fmt.Fprintf(w, err2.Error())
		}
	}

	fmt.Println("Execution time on sequential: ", time.Since(start))

}
