package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	fmt.Println("Response Status: ", response.Status)
	// fmt.Println("Contenaaaaaaaaaaat length: ", response.Status)

	if response.Status != "200 OK" {
		panic("Error while getting response")
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)

	}
	fmt.Println("\nrow data: ", data)
	fmt.Println("\ndata", string(data))
}
