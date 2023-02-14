package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer learning.")

	var age int  // this say that a variable that holds integer type
	var ptr *int //this say that a pointer that holds integer type.

	fmt.Println("Value of pointeris ", ptr) // <nil>
	fmt.Println("Value of variable ", age)  // 0 value put

	//  more about declare pointer variablr

	// & --> reference of memory of variable
	myAge := 23
	var ptrNew = &myAge
	fmt.Println("value of the pointer: ", ptrNew)
	fmt.Println("value of actual ointer: ", *ptrNew)

	// more explore
	fmt.Println("value of myage: ", myAge)
	// fmt.Println("value of myage: ", *myAge)
	fmt.Println("value of myage: ", *&myAge, &myAge, &*&myAge)

	// moore stuff
	*ptrNew = *ptrNew * 2
	fmt.Println("new pointer: by pointer ", *ptrNew)

	fmt.Println("new pointer: by actual variable", myAge)

}
