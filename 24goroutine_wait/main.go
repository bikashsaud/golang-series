package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //create a new WaitGroup
func main() {
	fmt.Println("Starting go routine wait.")

	websites := []string{
		"https://go.dev",
		"https://facebook.com",
		"https://google.com",
	}

	for _, website := range websites {
		go getStatusCode(website)
		// fmt.Printf("%s Status code : %d", website, status)
		wg.Add(1) // add new goroutine to wait group

	}
	wg.Wait()

}

func getStatusCode(endpoint string) int {
	defer wg.Done() // call wg.Done() to indicate that it has completed
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("error to getting status code: ", err)
	}
	fmt.Printf("Website: %s get status code: %d \n", endpoint, res.StatusCode)
	return res.StatusCode
}
