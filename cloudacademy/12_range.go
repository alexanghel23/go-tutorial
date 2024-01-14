package main

import "fmt"

type player struct {
	id   int
	rank int
}

func main() {
	bytes := []int{67, 108, 111, 117, 100, 65, 99, 97, 100, 101, 109, 121}
	for idx, value := range bytes {
		fmt.Print(string(value))
		if len(bytes)-1 == idx {
			fmt.Println()
		}
	}

	team := map[string]player{"John": {3, 10}, "Bob": {14, 11}}
	fmt.Println(team)

	for key, value := range team {
		fmt.Printf("%s -> %d\n", key, value)
	}

	for _, value := range team {
		fmt.Println(value)
	}

	for key := range team {
		fmt.Println("key:", key)
	}
}

/*
Range can actually be used to iterate over not only arrays, but also slices and maps. We can see this in action here.

On line 11, I'm creating a byte slice of ints, containing some mystery numbers. Lines 12 through to 17, I then use the range keyword to iterate over the byte slice and then convert each number to its equivalent character value, with the final result being CloudAcademy after the iteration finishes.
Lines 22 to 24 demonstrate how to use range to iterate over a map. This particular example shows how to receive both the key and value for each item returned during the iteration.

The next example, spanning lines 26 to 28, shows how to iterate over the same map, but this time using the underscore character to discard the key, since in this example, we're only interested in the value.
If we were to declare the key in the same way as the previous example on lines 22 to 24, but were not to reference it within the for loop, then the compiler would throw a declared and not used error, hence why the underscore is used.

The final example on lines 30 to 32 demonstrates the alternate requirement of working with just the key and not the value.

In summary, you have observed how to use the range keyword to support iterating over arrays, slices, and maps and how to use the underscore character to suppress compile time errors when not requiring to use the key within the for loop iteration.
*/
