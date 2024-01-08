package main

import "fmt"

var pl = fmt.Println

// varadic functions are functions that can receive an unknown number of values
// func funcName(parameters) returnType {Body}

func getSum2(nums ...int) int {
	pl(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	pl(arr)
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {

	vArr := []int{1, 2, 3, 7}
	pl(getArraySum(vArr))

}
