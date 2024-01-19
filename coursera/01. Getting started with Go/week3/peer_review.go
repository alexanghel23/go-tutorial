package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// func main() {
// 	slice := make([]int, 0, 3)

// 	for {
// 		fmt.Print("Enter an integer or press 'X' to exit): ")
// 		var input string
// 		fmt.Scanln(&input)

// 		if input == "X" || input == "x" {
// 			break
// 		}

// 		num, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Invalid input. Please enter a valid integer.")
// 			continue
// 		}

// 		slice = append(slice, num)
// 		sort.Ints(slice)

// 		fmt.Println("Sorted slice:", slice)
// 	}
// }

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numList := make([]int, 0, 3)
	var userInput string
	fmt.Println(numList)
	for {
		fmt.Print("Enter a value: ")
		fmt.Scan(&userInput)
		if strings.Compare(userInput, "x") == 0 {
			break
		}
		num, err := strconv.Atoi(userInput)
		if err != nil {
			continue
		}
		numList = append(numList, num)
		sort.Ints(numList)
		fmt.Println(numList)
	}
}
