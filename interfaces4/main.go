package main

import "fmt"

func describe(a interface{}) {
	fmt.Printf("Type = %T, value = %v\n", a, a)
}

func assert(i interface{}) {
	val, ok := i.(int)
	fmt.Println(val, ok)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am a int and my value is %d\n", i.(int))
	default:
		fmt.Printf("unknown type\n")
	}

}
func main() {
	i := 55
	describe(i)
	s := "HelloWorld"
	describe(s)
	strt := struct {
		Name string
	}{
		Name: "vamsikrishna",
	}

	describe(strt)

	assert(i)
	findType(i)
	assert(s)
	findType(s)
	findType(strt)

}
