package main

import (
	"fmt"
	simpleinterest "learn-package/simple-interest"
	"log"
)

var p, r, t = -5000.0, 10.0, 1.0

func init() {
	fmt.Println("Main package initialized.")
	if p < 0 {
		log.Fatal("Prinicipal is less than 0")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than 0")
	}
	if t < 0 {
		log.Fatal("Duration is less than 0")
	}
}

func main() {

	fmt.Println("Simple interst calculation.")
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interst is", si)
}
