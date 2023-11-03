package main

import "fmt"

type Income interface {
	source() string
	calculate() int
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

func (f FixedBilling) source() string {
	return f.projectName
}

func (t TimeAndMaterial) source() string {
	return t.projectName
}

func (f FixedBilling) calculate() int {
	return f.biddedAmount
}

func (t TimeAndMaterial) calculate() int {
	return t.hourlyRate * t.noOfHours
}

func main() {
	project1 := &FixedBilling{
		projectName:  "proj1",
		biddedAmount: 230,
	}

	project2 := &TimeAndMaterial{
		projectName: "proj2",
		noOfHours:   3,
		hourlyRate:  4,
	}
	project3 := &Advertisement{
		adName:     "proj3",
		CPC:        2,
		noOfClicks: 4,
	}

	incomeStreams := []Income{project1, project2, project3}
	sum := 0
	for _, v := range incomeStreams {
		sum = sum + v.calculate()
	}

	fmt.Println(sum)
}
