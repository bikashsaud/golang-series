package main

import "fmt"

func main() {
	fmt.Println("day5: Slicing ")
	var courses = []string{"math", "science", "Economics", "Social studies", "Business", "Education", "Computer Science"}
	fmt.Println(courses)

	var index int = 4
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
