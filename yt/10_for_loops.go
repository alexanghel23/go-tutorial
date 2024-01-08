package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	// for initialization; condition; postStatement {BODY}

	for x := 1; x <= 5; x++ {
		pl(x)
	}

	pl("\n")

	for x := 5; x >= 1; x-- {
		pl(x)
	}

}
