package main

import (
	"fmt"
)

var pl = fmt.Println

/*
multiline comment
*/

func main() {

	aNums := []int{1, 2, 3}

	for _, num := range aNums { // for _ means we are not interested in the index, only the value
		pl(num)
	}

	// //

	aStr := []string{"abc", "def", "ghi"}

	for i, str := range aStr { // for _ means we are not interested in the index, only the value
		pl(i, str)
	}

}
