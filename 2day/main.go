package main

import "fmt"

func main() {
	var username string = "Bikash saud."
	fmt.Println("variable type: %T \n ", username)

	var isAdmin bool = false
	fmt.Println("variable type: %T \n ", isAdmin)

	var number int = 12
	fmt.Println(number)
	fmt.Println("number : %T \n ", number)

	// no var style declaration

	salary := 12000.12
	//this is valorous operator used inside function
	fmt.Println(salary)

// 	check tyes
    // boolean
    var4 := true

    // shorthand string array declaration
    var5 := []string{"foo", "bar", "baz"}

    // map is reference datatype
    var6 := map[int]string{100: "Ana", 101: "Lisa", 102: "Rob"}

    // complex64 and complex128
    // is basic datatype
    var7 := complex(9, 15)

    // using %T format specifier to
    // determine the datatype of the variables

    fmt.Println("Using Percent T with Printf")
    fmt.Println()
    fmt.Printf("var4 = %T\n", var4)
    fmt.Printf("var5 = %T\n", var5)
    fmt.Printf("var6 = %T\n", var6)
    fmt.Printf("var7 = %T\n", var7)
}
