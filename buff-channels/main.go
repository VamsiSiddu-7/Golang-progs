package main

import (
	"fmt"
	"sync"
	"time"
)

// func write(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("successfully wrote", i, "to ch")
// 	}
// 	close(ch)
// }

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Println("ended goroutine ", i)
	wg.Done()
}

func main() {
	// ch := make(chan int, 2)
	// go write(ch)
	// time.Sleep(2 * time.Second)
	// for v := range ch {
	// 	fmt.Println("read value", v, "from ch")
	// 	time.Sleep(2 * time.Second)
	// }
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		go process(i, &wg)
		wg.Add(1)
	}
	wg.Wait()
}
