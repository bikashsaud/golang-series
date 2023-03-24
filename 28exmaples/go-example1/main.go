package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Example 1: Thhis example show how to change modify struct type by using its pointer address through but if we try modift struct type is not possible")

	p1 := Person{Name: "Bikash Saud", Age: 12}
	fmt.Printf("Person p1 Name: %s and Age: %d", p1.Name, p1.Age)
	fmt.Printf("\n")

	p2 := &Person{Name: "Harish Saud", Age: 32}
	fmt.Printf("Person p2 Name %s and Age: %d", p2.Name, p2.Age)
	fmt.Printf("\n")

	modifyPerson(p1)
	fmt.Println("Person p1", p1.Name, p1.Age)
	fmt.Printf("\n")

	modifyPersonPtr(p2)
	fmt.Println("Person p2", p2.Name, p2.Age)

}

func modifyPerson(p Person) {
	p.Name = "Harmanprit Singh"
	p.Age = 21
	fmt.Println("Inside the function is applied but not globle moidfy.", p)
}

func modifyPersonPtr(p *Person) {
	p.Name = "Bikash Said"
	p.Age = 21
}


