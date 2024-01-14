package main

import "fmt"

type player struct {
	id   int
	rank int
}

func main() {
	map1 := make(map[string]string)
	map1["key1"] = "value1"
	map1["key2"] = "value2"
	map1["key3"] = "value3"
	fmt.Println(map1)

	value1 := map1["key1"]
	fmt.Println(value1)

	//map1["key1"] = 10 //compile error
	map1["key1"] = "value1.1"
	delete(map1, "key2")
	map1["newkey"] = "value4"
	fmt.Println(map1)

	team := map[string]player{"John": {3, 10}, "Bob": {14, 11}}

	fmt.Println(team)
}

/*
Go provides a dynamic growable, associative data type called a map.

A map is a collection of key-value pairs and behaves much like a dictionary does, in that when you query a map with a key you get back the value that was mapped to that key.

Let's see how this works.
On line 11, map1 is defined as a map of string to string. In this case, both the key and value are typed as strings.
Lines 12 to 14 then add three key-value pairs to the map.
Retrieving a value back out of the map is done by providing the key name within square brackets as per line 17.
You have to remember that maps are typed at compile time and you cannot mix and match types after the map has been declared. For example, line 20 would cause a compile time error since the value attempting to be set is of type int, but the map was originally typed with both the key and values both being of type string.

A slightly more advanced version of a map is implemented on line 26.
Here I'm declaring a team map. The player's name is the key. Therefore the key is typed as a string. The value this time is using a custom struct type. The player struct type defined on lines 5 to 8. The entire team map is initialized and loaded on a single line using the curly bracket notation.

In summary, you've just observed how to use the make function to create and declare a variable of type map, how to retrieve values out of the map by specifying their key, creating and initializing a map in one hit on a single line, and using a custom struct as the data type for the maps value part.
*/
