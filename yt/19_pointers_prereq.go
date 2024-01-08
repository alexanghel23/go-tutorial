package main

import "fmt"

var pl = fmt.Println

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func main() {

	f3 := 5
	pl("f3 before func: ", f3)

	// function won't change the value of f3
	changeVal(f3)
	pl("f3 after func:", f3)

}
