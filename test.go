package main

import (
	"io"
	"log"
	"net/http"
)

type testStruct struct {
	responses map[int]string
}

func main() {

	client := http.Client{}

	req, err := client.Get("https://catalog.dedeman.ro/api/live/list")

	if err != nil {
		panic(err)
	}

	//res, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)

}
