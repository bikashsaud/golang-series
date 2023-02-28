package main

import "fmt"

func main() {

	defer fmt.Println("DEFERed statement.")

	println("Hello world!")

	defer fmt.Println("Start from here...")

	deferCall()
}

func deferCall(){
	for i := 0; i <10; i++{
		defer fmt.Println("Defer call value: ", i)
	}
}

// Hello world, i value, start from here ..., DEEFRed stamtents