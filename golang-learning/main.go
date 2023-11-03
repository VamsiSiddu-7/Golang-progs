package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("Hello world")
	}

	a()
	fmt.Println(area(2, 3))
}

func area(a, b int) (cal int) {
	cal = a * b
	return
}
