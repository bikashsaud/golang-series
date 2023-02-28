package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("if else control statements...")
	reader := bufio.NewReader(os.Stdin)

	// logincount := 12
	println("Enter user count:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading user input: ", err)
	} else {

		logincount, _ := strconv.Atoi(input)
		var result string

		if logincount < 10 {
			result = "SImple User"
		} else if logincount > 10 {
			result = "Most valuable user."
		} else {
			result = "Fine User/"
		}
		fmt.Println(result, "logincount: ", logincount%2)

		if logincount%2 == 0 {
			print("input number is even")
		} else if logincount == 0 {
			print("input number is not even and nor odd.")
		} else {
			print("Input number is odd.")
		}
	}

}
