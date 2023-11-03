package main

import "fmt"

type VolwelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("vamsikrishna")
	var v VolwelsFinder = name
	vs := v.FindVowels()
	for i := 0; i < len(vs); i++ {
		fmt.Printf("%c", vs[i])
	}
	fmt.Println()

}
