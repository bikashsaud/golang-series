package main

import "fmt"

func main() {
	fmt.Println("Exportable or not exportable methods")
	college := College{
		Name:     "Amrit Science Campus",
		Address:  "Lainchaur Kathmandu",
		Students: 1200,
		Teachers: 70}

	newStudents := 120
	leaveThisYear := 35
	totalStudents, _ := college.SetStudentNumber(newStudents, leaveThisYear)
	fmt.Println("Total students: ", totalStudents)

	// getter
	teachers, _ := college.getTeacherNumber()
	fmt.Println("GET TOTAL TEACHER: ", teachers)
}

type College struct {
	Name     string
	Address  string
	Students int
	Teachers int
}

// works as getter/setter
func (c College) SetStudentNumber(newS int, oldLeave int) (int, error) {
	return c.Students + newS - oldLeave, nil
}

// This is not exportable method noe
func (c College) getTeacherNumber() (int, error) {
	return c.Teachers, nil
}

