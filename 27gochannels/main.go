package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("GO Channels ")

	myCh := make(chan int, 2) // 2 is buffer for channel run
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// write only here.
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 10
		close(myCh) //close the channel after complete done
		wg.Done()
	}(myCh, wg)

	// Read only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChannelOpen := <-myCh // isChannelOpen says is channel mis close and val is returned val
		fmt.Println(val, isChannelOpen)

		// val1, isChannelOpen := <-myCh // isChannelOpen says is channel mis close and val is returned val

		// fmt.Println(val1, isChannelOpen)
		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
