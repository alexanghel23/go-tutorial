package main

import "fmt"

var pl = fmt.Println

func changeVal2(myPtr *int) { // * symbol allows us to access the place in memory
	*myPtr = 12
}

func main() {
	f4 := 10
	var f4Ptr *int = &f4
	pl("f4 address: ", f4Ptr)
	pl("f4 value: ", *f4Ptr)

	pl("f1 before func: ", f4)

	// acces the specific location in memory by adding address of the operator &
	changeVal2(&f4)
	pl("f4 after func:", f4)

}
