package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	ID    int
	basic int
	pf    int
}

type Contractor struct {
	ID    int
	basic int
}

type FreeLancer struct {
	ID          int
	ratePerHour int
	totalHours  int
}

func (p Permanent) CalculateSalary() int {
	return p.basic + p.pf
}

func (c Contractor) CalculateSalary() int {
	return c.basic
}

func (f FreeLancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func main() {
	e1 := Permanent{
		ID:    1,
		basic: 100,
		pf:    10,
	}

	e2 := Contractor{
		ID:    2,
		basic: 100,
	}

	e3 := FreeLancer{
		ratePerHour: 2,
		totalHours:  3,
	}

	employees := []SalaryCalculator{e1, e2, e3}
	totalSalary := 0
	for _, e := range employees {
		totalSalary = totalSalary + e.CalculateSalary()
	}

	fmt.Println(totalSalary)
}
