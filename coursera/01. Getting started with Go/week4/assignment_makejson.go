package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

// In Go, when you define struct fields with lowercase names, they are considered unexported, meaning they won't be accessible outside the package.
// To make them accessible for JSON marshaling, you should use uppercase initial letters for the struct fields.
type Person struct {
	Name    string
	Address string
}

func main() {

	fmt.Println("Enter name:")

	name_scan := bufio.NewScanner(os.Stdin)
	name_scan.Scan()

	name := name_scan.Text()

	fmt.Println("Enter address:")

	address_scan := bufio.NewScanner(os.Stdin)
	address_scan.Scan()

	address := address_scan.Text()

	p1 := Person{name, address}

	js, _ := json.Marshal(p1)

	fmt.Println(string(js))

}
