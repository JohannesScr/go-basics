// Package mymath implements functions to do math
package mymath

import "fmt"

// Sum take unlimited integers and adds them all together and returns
// the total
func Sum(x ...int) int {
	total := 0
	for x := range x {
		total += x
	}
	return total
}

// SumEven take unlimited integers and adds only the even numbers together
// and returns the total
func SumEven(x ...int) int {
	total := 0
	for x := range x {
		total += x
	}
	return total
}

// noExportedFunc will not show up in documentation. Only exported functions that
// start with a capital letter such as ExportedFunc.
func noExportedFunc() {
	fmt.Println("I am not exported")
}
