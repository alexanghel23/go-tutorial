package main

import (
	"fmt"
)

//const maxValue //compile error

/*
constants are scoped
they only exist within the scope that they were declared within
constants are typed
constants must have a defined value at the time of declaration
constants cannot be reassigned anytime after the initial declaration.
*/

const population int = 10000

func main() {
	const name = "CloudAcademy"

	//name = "Blah" //compile error

	if true {
		const color = "red"

		fmt.Println(population)
		fmt.Println(name)
	}

	//fmt.Println(color) //compile error
}
