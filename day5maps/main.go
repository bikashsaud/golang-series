package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps GO...")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["py"] = "python"
	languages["jv"] = "java"
	languages["rb"] = "ruby"
	languages["php"] = "php"
	languages["go"] = "golang"

	fmt.Println("Languages: ", languages)
	fmt.Println(languages["jv"])

	// delete

	delete(languages, "jv")
	fmt.Println("languages after delete: ", languages)

	// loops in go for maps
	for key, val := range languages {
		fmt.Println("key: ", key)
		fmt.Println("val: ", val)

	}
}
