package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.

The program should be written as a loop.
Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.

The program adds the integer to the slice, sorts the slice,
and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any
number of integers which the user decides to enter.

The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

var pl = fmt.Println

func main() {

	var sli = make([]int, 3)

	for true {
		fmt.Println("Enter integer:")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		ipt := scanner.Text()

		iptc, err := strconv.Atoi(ipt)

		if err != nil {

			if ipt == "X" {
				pl("Exiting loop")
				break
			}
			pl("Error, not an integer!")
			continue
		} else {
			sli = append(sli, iptc)
			pl(sli)
		}
	}
	pl(sli)
}
