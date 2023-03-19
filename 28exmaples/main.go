package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Maths in golang...")

	// var numOne int = 2
	// var numTwo float64 = 4

	// fmt.Println("the sum of numbers: %f", numOne+numTwo)

	// random numbers

	// fmt.Println(rand.Intn(2))

	// crypto random

	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRandNum)

}
