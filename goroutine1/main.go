package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello goroutine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello goroutine awake and going to write to done")
	done <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dchnl := make(chan int)
	go digits(number, dchnl)
	for digit := range dchnl {
		sum = sum + (digit * digit)
	}
	squareop <- sum
}

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number = number / 10
	}
	close(dchnl)
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dchnl := make(chan int)
	go digits(number, dchnl)
	for digit := range dchnl {
		sum = sum + (digit * digit * digit)
	}

	cubeop <- sum
}
func main() {
	done := make(chan bool)
	number := 589
	squarech := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, squarech)
	go calcCubes(number, cubech)
	squares, cubes := <-squarech, <-cubech
	fmt.Println(squares, cubes)
	fmt.Println("Main goroutine to call hello go routine")
	go hello(done)
	<-done
	fmt.Println("main recieved data....")
}
