package main

import "fmt"

func main() {
	s := "Señor"
	var bs = []rune(s)
	fmt.Println(bs)

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%c", bs[i])
	}
	fmt.Println()
}
