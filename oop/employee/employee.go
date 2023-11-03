package employee

import "fmt"

type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

func New(firstName string, lastName string, totalLeaves int, leavestaken int) employee {
	e := employee{firstName: firstName, lastName: lastName, totalLeaves: totalLeaves, leavesTaken: leavestaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
