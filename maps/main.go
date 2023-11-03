package main

import "fmt"

func main() {
	var es map[string]int
	employeeSalary := map[string]int{
		"John":  1,
		"jaffa": 2,
		"Daffa": 3,
		"Paffa": 4,
	}
	fmt.Printf("%#v %#v\n", employeeSalary, es)
	fmt.Printf("%#v\n", employeeSalary["kaffa"])
	value, ok := employeeSalary["dippam"]

	fmt.Println(value, ok)

	for key, value := range employeeSalary {
		fmt.Println("key,value", key, value)
	}

	delete(employeeSalary, "John")

	fmt.Println(employeeSalary)

}
