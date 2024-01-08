package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// in terminal run go build -> a file called 24_command_line is created
	// in terminal run ./24_command_line 24 54 66 1 -> and array is printed
	fmt.Println(os.Args)

	// remove file name from the array
	args := os.Args[1:]

	var iArgs = []int{}

	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max value is ", max)

}
