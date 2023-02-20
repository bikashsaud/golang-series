package main

import "fmt"

func main() {
	fmt.Println("Struct in GO.")
	// no inheretance in go lang; No Super no parent class

	bikash := User{"BIkash Saud", "saud@gmail.com", true, 23}
	fmt.Println(bikash)
	fmt.Printf("bikash detail : %+v\n", bikash)
	fmt.Printf("Nsme: %v Email: %v", bikash.Name, bikash.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// struct name is always in title form
// variable name is also Title case
