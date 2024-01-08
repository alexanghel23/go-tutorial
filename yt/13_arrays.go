package main

import (
	"fmt"
)

var pl = fmt.Println

/*
multiline comment
*/

func main() {

	var arr1 [5]int // 5 is the size of the array; it will be filled out with defaults (0 for integers, 0.0 for float, false for boolean, "" for strings)
	arr1[0] = 1
	pl(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	pl(arr2[0])

	arr21 := []int{1, 10, 100} // the size of the array can be left out
	pl(arr21)

	pl("Array length is ", len(arr2))

	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	// multidimensional array

	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	// pl(arr3)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

}
