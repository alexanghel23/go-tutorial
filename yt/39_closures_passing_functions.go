package main

import (
	"fmt"
)

var pl = fmt.Println

// functions that don't have to be associated with an identifier but very often they are associated with a variable

// func changeVarr(samp int) int {
// 	samp += 1
// 	return samp
// }

func useFunc(f func(int, int) int, x, y int) {
	pl("Answer:", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
	intSum := func(x, y int) int { return x + y }
	pl("5+4=", intSum(5, 4))

	samp1 := 1
	changeVar := func() {
		samp1 += 1
	}

	changeVar()
	// changeVarr(samp1)

	pl("samp1=", samp1)

	useFunc(sumVals, 5, 8)

}
