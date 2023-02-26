package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("switch  and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("1, OPEN NEW.")
		fallthrough
	case 2:
		fmt.Println("2, MOVE 2 step forward.")
	case 3:
		fmt.Println("3, MOVE 3 step forward.")
	case 4:
		fmt.Println("4, MOVE 2 step forward.")
	case 5:
		fmt.Println("5, MOVE 5 step forward.")
	case 6:
		fmt.Println("6, MOVE 6 step forward.")
	default:
		fmt.Println("try again.")

	}

}
