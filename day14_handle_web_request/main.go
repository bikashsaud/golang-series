package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("Web request response handle in Go.")
	response, err := http.Get(url)
	checkNilError(err)
	fmt.Println("Reponse handle in Go.", response)

	defer response.Body.Close() // its callers responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	fmt.Println(string(databytes), 23400)
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
