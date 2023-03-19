package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Go Race conditions")
	wg := &sync.WaitGroup{}
	mute := &sync.RWMutex{} // mute := &sync.Mutex{} also use as previous

	var score = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mute.Lock()
		score = append(score, 1)
		mute.Unlock()
		wg.Done()
	}(wg, mute)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mute.Lock()
		score = append(score, 2)
		mute.Unlock()

		wg.Done()

	}(wg, mute)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mute.Lock()

		score = append(score, 3)
		mute.Unlock()

		wg.Done()

	}(wg, mute)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mute.RLock()
		fmt.Println(score)
		mute.RUnlock()
		wg.Done()

	}(wg, mute)

	wg.Wait()
	fmt.Println("Scores:", score)

}
