package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello from bikash")
	splitVar := strings.SplitAfter("from", "bikash")
	fmt.Println(splitVar)
}
