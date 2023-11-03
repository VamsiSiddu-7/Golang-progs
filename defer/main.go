package main

import "fmt"

func main() {
	name := "naveen"

	for _, v := range name {
		defer fmt.Printf("%c", v)
	}

}
