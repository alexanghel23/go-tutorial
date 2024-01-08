package main

import (
	"fmt"
)

var pl = fmt.Println

type Tsp float64 //tea spoon
type TBs float64 // table spoon
type ML float64  // mililiters

// Next sections is the preferred way to do it

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBTtoML(tsp TBs) ML {
	return ML(tsp * 14.79)
}

func main() {

	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsps = %.2f\n", ml1)

	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBS = %.2f\n", ml2)

	pl("2 tsp + 4 tsp =", Tsp(2), Tsp(4))
	pl("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))

}
