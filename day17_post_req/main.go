package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post request for  go language...")
	HandlePostRequest()

}

func HandlePostRequest() {
	const url = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
		{
			"courseName": "Lets go with golang.",
			"price":0,
			"platform": "Youtube.com"
		}
	`)
	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response: ", string(content))
}
