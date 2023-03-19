package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup                   //create a new WaitGroup
var signals = []string{"test websites"} //pointer
var mut sync.Mutex                      //pointer

func main() {
	fmt.Println("Starting go routine wait.")

	websites := []string{
		"https://mobile.twitter.com",
		"https://psc.gov.np/",
		"https://go.dev",
		"https://facebook.com",
		"https://google.com",
		"https://identeq.treeleaf.ai",
	}

	for _, website := range websites {
		fmt.Printf("Request:--> website: %v request time: %s\n", website, time.Now().UTC())
		go getStatusCode(website)
		// fmt.Printf("%s Status code : %d", website, status)
		wg.Add(1) // add new goroutine to wait group

	}
	wg.Wait()
	fmt.Println(signals)

}

func getStatusCode(endpoint string) int {
	defer wg.Done() // call wg.Done() to indicate that it has completed
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("error to getting status code: ", err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		// fmt.Printf("Website: %s get status code: %d \n", endpoint, res.StatusCode)
		fmt.Printf("Response--> Website: %s, end time: %s  CODE: %d\n", endpoint, time.Now().UTC(), res.StatusCode)
		return res.StatusCode
	}
	return res.StatusCode
}
