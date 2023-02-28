package main

import "fmt"

func main() {
	fmt.Println("Start function in golang.")
	greetings()

	age := ageCalculator(2053, 2079)
	fmt.Println("AGE: ", age)

	fmt.Println("ARGS ADDER RESULT")

	allAdder := argsAdder(12, 23, 42, 45)
	fmt.Println("All added sum = ", allAdder)

}

func greetings() {
	fmt.Println("Hello world, Good morning")
}

func ageCalculator(birthYear int, currentYear int) int {
	fmt.Println("Adding")
	return currentYear-birthYear  
}

func argsAdder(args ...int) int {
	total := 0
	for _, arg := range args {
		total += int(arg)
	}
	fmt.Println("ARGS ADDER")
	return total

}
