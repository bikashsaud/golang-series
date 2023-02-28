package main

import "fmt"

func main() {
	fmt.Println("For loop")

	subjects := []string{"nepali", "english", "social", "Science"}

	// forloop
	for i := 0; i < len(subjects); i++ {
		fmt.Println("LOOP value: ", subjects[i])
	}

	for index, sunject := range subjects {
		fmt.Println("Index: ", index, "Subject: ", sunject)
	}

	// While loop: there is bo dedicated while loop in GOlang
	fmt.Println("FOR LOOP...")
	k := 0
	for k < len(subjects) {
		fmt.Println(k, subjects[k])
		k++
	}

	// DO-WHILE LOOP

	j := 0
	fmt.Println("DO-WHILE LOOP...")
	for {
		fmt.Println("Index: ", j, "Subject: ", subjects[j])
		j++
		if j >= len(subjects) {
			break
		}
	}

}
