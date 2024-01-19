package main

import (
	"encoding/json"
	"fmt"
)

/*

JSON Properties
- All unicode, human readable
- Fairly compact representation
- Types can be combined recursively: array of structs, struct in struct etc.
*/

type Person struct {
	name  string
	addr  string
	phone string
}

func main() {

	p1 := Person{name: "joe", addr: "street", phone: "123"}

	// JSON Marshalling - Generating json representation from a go object
	barr, _ := json.Marshal(p1) // Marshal returns json representation as []byte

	fmt.Println(barr)

	var p2 Person

	err := json.Unmarshal(barr, &p2) // Unmarshal() converts a json []byte into a Go object; Pointer to Go object is passed

	fmt.Println(p2, err)

}
