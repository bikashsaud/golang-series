package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://lco.dev:3000/learn?courcename=reactjs&page=1"

func main() {
	fmt.Println("URL Handle")
	fmt.Println(myurl)
	// parsingUrl
	response, _ := url.Parse(myurl)

	fmt.Println(response.Scheme)
	fmt.Println(response.Host)
	fmt.Println(response.Path)
	fmt.Println(response.RawQuery)
	fmt.Println(response.Port())

	queryParams := response.Query()
	fmt.Println(queryParams)

	for key, val := range queryParams {
		fmt.Println("val: ", val)
		fmt.Println("Key: ", key)
	}

	// parts of url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=bikash",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
