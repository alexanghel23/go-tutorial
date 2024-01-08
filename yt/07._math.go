package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {

	pl("5 + 4 = ", 5+4)
	pl("5 - 4 = ", 5-4)
	pl("5 * 4 = ", 5*4)
	pl("5 / 4 = ", 5/4)
	pl("5 % 4 = ", 5%4)

	// mInt := 1
	// mInt = mInt + 1

	// mInt += 1
	// mInt++

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)

	randNum := rand.Intn(50) + 1

	pl("Random: ", randNum)

	pl("Random: ", randNum)

}
