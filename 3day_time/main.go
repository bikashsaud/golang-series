package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome GO Time...")
	now_time := time.Now().Format("01-02-2006 11:04:05 Monday")
	fmt.Println(now_time)
	unique_time := time.Now().UnixMicro()
	fmt.Println("Unique time stamp: ", unique_time)
	// add date
	t := time.Date(1996, time.July, 20, 19, 30, 0, 0, time.UTC)
	fmt.Println("MY date of birth is : ", t)
}
