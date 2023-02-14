package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("User Login")
	fmt.Println("Please enter your username and password.")

	buf := bufio.NewReader(os.Stdin)
	username, _ := buf.ReadString('\n')

	fmt.Println("your username is: ", username)

	validate_username, err := strconv.ParseFloat(strings.TrimSpace(username), 64)

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("username: %v\n type: %T\n", validate_username, validate_username)
	}
}
