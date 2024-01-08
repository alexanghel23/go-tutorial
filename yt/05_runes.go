package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {

	rStr := "abcdefg"

	pl("Rune count: ", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

	// Creating a rune
	// In Go language, a Rune Literal is expressed as one or more characters enclosed in single quotes

	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	// Displaying rune and its type
	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s", rune1,
		rune1, reflect.TypeOf(rune1))

	fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %s", rune2,
		rune2, reflect.TypeOf(rune2))

	fmt.Printf("\nRune 3: Unicode: %U; Type: %s", rune3,
		reflect.TypeOf(rune3))

}
