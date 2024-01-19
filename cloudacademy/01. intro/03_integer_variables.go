package main

import (
	"fmt"
)

func main() {
	var count1 int
	var count2 int = 100
	count3 := 200
	count4 := count2 + count3

	fmt.Println(count1) //prints 0
	count1 = 10
	fmt.Println(count1) //prints 10

	fmt.Println(count2)
	fmt.Println(count3)
	fmt.Println(count4)

	count3++
	count4--

	fmt.Println(count3)
	fmt.Println(count4)
}
