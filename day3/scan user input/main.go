package main

import "fmt"
import "os"
import "bufio"

func main() {
    welcome:= "welcome to user package code. "
    fmt.Println(welcome)
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter your username: ")
    // comma Ok // err ok
    username, _ := reader.ReadString('\n')
    fmt.Println("Enter your password: ")
    password, _ := reader.ReadString('\n')

    fmt.Println("Your Username: ", username)
    fmt.Println("Your Password: ", password)

//     types

fmt.Printf("Type of your Input username: %T \n password: %T^ \n ", username, password)
}