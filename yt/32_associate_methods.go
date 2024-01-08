package main

import (
	"fmt"
)

// This is the way to do it, not the previous one

var pl = fmt.Println

type Tsp float64 //tea spoon
type TBs float64 // table spoon
type ML float64  // mililiters

func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}

func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

func main() {

	tsp1 := Tsp(4)
	fmt.Printf("%.2f tsp = %.2fmL\n", tsp1, tsp1.ToMLs())
}
