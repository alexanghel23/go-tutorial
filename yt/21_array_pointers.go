package main

import "fmt"

var pl = fmt.Println

func dblArrVals(arr *[4]int) { // * symbol allows us to access the place in memory
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))

	for _, val := range nums {
		sum += val
	}

	return (sum / numSize)
}

func main() {

	pArr := [4]int{1, 2, 3, 4}

	dblArrVals(&pArr)

	pl(pArr)

	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average: %.3f\n", getAverage(iSlice...))

}
