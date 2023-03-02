package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Post form data to server.")

	PostFormRequest()

}

func PostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	// form_data
	data := url.Values{}
	data.Add("first_name", "Bikash")
	data.Add("last_name", "Saud")
	data.Add("email", "saud@gmail.com")
	data.Add("address", "test address")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // Close the connection request
	// fmt.Println(response.Body, 999999999)
	// read response body
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
