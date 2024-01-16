package main

import "fmt"

/*


/*
Maps
- they are golang's implementation of a hash table
- use make() to create a map
- or you can define one with a map literal

Accessing maps
- referencing a value with [key]

Map functions
- two-value assignment tests for existence of the key
   id, p := idMap["joe"], where is a boolean equal to the existence of the key "joe" in the map
- len: how many key-value pairs are in the map

Iterating through map
- use a for loop with a range keyword
- two-value assignment with range

*/

var idMap map[string]int

func main() {

	idMap = make(map[string]int)

	idMap2 := map[string]int{"joe": 123}

	fmt.Println(idMap, idMap2)
	fmt.Println(idMap2["joe"]) // returns zero if key is not present
	idMap2["jane"] = 456
	idMap2["gina"] = 789
	delete(idMap2, "joe")

	fmt.Println(idMap2)

	_, p := idMap2["jane"]
	fmt.Println(p)

	for key, value := range idMap2 {
		fmt.Println(key, value)
	}

}
