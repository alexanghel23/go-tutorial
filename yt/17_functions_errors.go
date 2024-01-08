package main

import "fmt"

var pl = fmt.Println

// func funcName(parameters) returnType {Body}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x / y, err
	}
}

func main() {

	pl(getQuotient(5, 0))

}
