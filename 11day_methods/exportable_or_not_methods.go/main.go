package main

import "fmt"

func main() {
	fmt.Println("Exportable or not exportable methods")
	college := College{
		Name:     "Amrit Science Campus",
		Address:  "Lainchaur Kathmandu",
		Students: 1200,
		Teachers: 70}

	college.COllegeDetail()
	newStudents := 120
	leaveThisYear := 35
	totalStudents, _ := college.GetStudentNumber(newStudents, leaveThisYear)
	fmt.Println("Total students: ", totalStudents)

}

type College struct {
	Name     string
	Address  string
	Students int
	Teachers int
}

// Exportable methond -> first later is capital
func (c *College) COllegeDetail() {
	fmt.Println("COllege Detail")
	fmt.Println("COllege: ", c.Name, c.Address)
	fmt.Println("Toatl Students: ", c.Students)
	fmt.Println("Total Teachers: ", c.Teachers)
}

// works as getter
func (c College) GetStudentNumber(newS int, oldLeave int) (int, error) {
	return c.Students + newS - oldLeave, nil
}
