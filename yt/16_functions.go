package main

import "fmt"

var pl = fmt.Println

// func funcName(parameters) returnType {Body}

func sayHello() {
	pl("hello")
}

func getSum(x int, y int) int {
	return x + y
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func main() {

	// sayHello()

	pl(getSum(5, 4))
	pl(getTwo(8))

}
