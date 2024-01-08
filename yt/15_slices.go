package main

import (
	"fmt"
)

var pl = fmt.Println

/*
multiline comment
*/

func main() {

	// var name []datetype

	sl1 := make([]string, 6)

	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	pl("Slice size: ", len(sl1))

	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	for _, x := range sl1 {
		pl(x)
	}

	sArr := [5]int{1, 2, 3, 4, 5}

	sl3 := sArr[0:2]

	pl("First 3 values: ", sArr[:3])
	pl("Last 3 values: ", sArr[2:])

	// Chaning the array will also change the slice
	sArr[0] = 10
	pl("sl3", sl3)

	// Interestingly, chaning the slice will also change the array
	sl3[0] = 7
	pl("sArr", sArr)

	sl3 = append(sl3, 12)
	pl("sl3", sl3)
	pl("sArr", sArr)

	sl4 := make([]string, 6)

	pl("sl4: ", sl4)
	pl("sl4[0]: ", sl4[0]) //shows a value of nil

}
