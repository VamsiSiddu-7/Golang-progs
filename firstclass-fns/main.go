package main

import "fmt"

type add func(a int, b int) int

func main() {
	func(n string) {
		fmt.Println(n)
	}("hello world")

	var a add = func(a int, b int) int {
		return a + b
	}

	s := a(5, 6)
	fmt.Println("Sum", s)
	k := simple()
	fmt.Println(k(100, 230))

	c := 5

	func() {
		fmt.Println("c == ", c)
	}()

	var f int = 1

	switch {
	case f == 1:
		fmt.Println("1")
		fallthrough
	case f == 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

}

func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}
