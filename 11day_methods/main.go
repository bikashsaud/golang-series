package main

import "fmt"


func main() {
	fmt.Println("METHODS in GOLANG: ")
	rect := Rectangle{width: 3, height: 4}
    fmt.Println("area:", rect.area())


}

type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

