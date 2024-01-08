package main

import (
	"fmt"
)

var pl = fmt.Println

/*
multiline comment
*/

func main() {

	aStr1 := "abcde"
	rArr := []rune(aStr1) //converts the string into an array of runes

	pl(rArr)

	for _, v := range rArr {
		fmt.Printf("Rune array: %d\n", v)
	}

	byteArr := []byte{'a', 'b', 'c'}
	pl(byteArr)
	bStr := string(byteArr[:])
	pl("I'm a string: ", bStr)

}
