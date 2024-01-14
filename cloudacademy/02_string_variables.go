package main

import "fmt"

// global variables, not just in the current file but anywhere in the wider application since they have been declared outside the main function and they names start with capital letter
var (
	Name1 string = "global1"
	Name2 string = "global2"
)

// global scope but only in the current file since their names start with lowercase character
var var1, var2 string = "local1", "local2"

func main() {

	var name1 string // the compiler will assing the variable with zero value; for strings that will be the empty string; its scope is limited to the main function; when the main function completes this variable and all other variables declared in the function will cease to exist
	var name2 string = "CloudAcademy"
	name3 := "Devops 2024"
	name4 := name2 + " " + name3

	fmt.Println(name1)
	name1 = "Blah"
	fmt.Println(name1)

	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)

	fmt.Printf("%v -- %v\n", name2, name3)

	fmt.Println(Name1, Name2, var1, var2)

}
