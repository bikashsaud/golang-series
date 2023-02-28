package main

import "fmt"

func main() {
	fmt.Println("BREAK CONDIOTION")

	for i := 0; i < 10; i++ {
		fmt.Println("value: ", i)
		if i == 6 {
			break
		}
	}

	fmt.Println("CONTINUE...")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue

		}
		fmt.Println("continue value: ", i)

	}

	fmt.Println("CONTINUE ANOTHER WAY...")
	l := 0
	for l < 10 {
		if l == 3 {
			l++
			continue
		}
		fmt.Println("loop value: ", l)
		l++
	}

	fmt.Println("JUMP CONDITIONS...")

	for i:=0; i<10; i++ {
		fmt.Println("jump conditions: ", i)
		if i==3{
			goto command
		}

	}
	command:
		fmt.Println("YOu must improve you loop code.")
}

