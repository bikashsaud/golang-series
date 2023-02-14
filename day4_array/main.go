package main

import "fmt"

func main() {
	fmt.Println("Go ARRAY Tutorial")

	var students [4]string

	students[0] = "bikash"
	students[3] = "Rahul"
	// students[1] = "Ram";
	// students[2] = "Sita";

	fmt.Println("students:", students)

	// another way

	marks := [6]int{1, 2, 3, 4, 5}

	fmt.Println("new marks array: ", marks)

	// another way

	var newMarks [6]string
	newMarks[0] = "new marks1"
	newMarks[5] = "New marks 2"

	fmt.Println("new marks array: ", newMarks)

	// type check error

	// typechecker := [3] string{1,2,3,4}
	// fmt.Println(typechecker)

	// EG for outof index
	// tyeCheck := [3] string{"New marks", "New marks", "New marks", "New marks"}
	// fmt.Println(tyeCheck)

	// get length of array
	fmt.Println("length of newMark array: ", len(newMarks))

	// another way to initialised
	var studentsList = [3]string{"pravesh", "prem", "janak"}
	fmt.Println("student list: ", studentsList, "length stdents list:", len(studentsList))
}
