package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Get Request Handle.")

	PerformGetRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:3000/get"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.Status != "200 OK" {
		panic("Error while getting response")
	}

	// use strings builder...
	var responseString strings.Builder
	data, err := ioutil.ReadAll(response.Body)
	byteCounter, _ := responseString.Write(data)

	if err != nil {
		panic(err)
	}
	fmt.Println("\nrow data: ", byteCounter)
	fmt.Println("\ndata", string(responseString.String()))
	fmt.Println("\ndata", responseString)
	
}
