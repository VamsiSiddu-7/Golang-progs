package main

import "fmt"

func sendData(sendCh chan<- int) {
	sendCh <- 1
}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
func main() {
	ch := make(chan int)
	go sendData(ch)
	fmt.Println("Recieved data:", <-ch)

	ch1 := make(chan int)
	go producer(ch1)

	for k := range ch1 {
		fmt.Println("Recieved", k)
	}
}
