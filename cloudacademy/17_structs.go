package main

import "fmt"

type person struct {
	firstname string
	surname   string
}

type lecture struct {
	name       string
	instructor person
	duration   int //seconds
}

func main() {
	lectures := []lecture{
		lecture{"Structs", person{"Jeremy", "Cook"}, 300},
		lecture{"Pointers", person{"Jeremy", "Cook"}, 300},
		lecture{"Functions", person{"Jeremy", "Cook"}, 300},
	}

	for _, lecture := range lectures {
		name := lecture.name
		instructor := fmt.Sprintf("%s %s",
			lecture.instructor.firstname,
			lecture.instructor.surname)
		duration := lecture.duration

		fmt.Printf("Lecture: '%s', Author: %s, Duration: %d secs\n", name, instructor, duration)
	}

}

/*
When it comes to creating custom data types, Go provides the Struct keyword.
A Struct provides you with the capability of creating a typed collection of related fields.
They are used to create custom data types by combining multiple related attributes together.
Structs can be nested within each other, enabling you to build up custom types that have a hierarchy of related fields.

In the example presented here, two custom structs are declared.
The person struct, defined on lines five to eight, defines a person type abstraction which is used to represent a person's first name and surname.
The second struct type, spanning lines 10 to 14, defines a lecture data type abstraction, used to represent a lecture which has a given name,
is presented by an instructor, and has a duration.
The instructor field within the lecture struct on line 12 is typed using the previously declared person struct.
This demonstrates how to perform composition using structs.

Later on within the main function, spanning lines 18, 19, and 20, three lecture literals are created using curly bracket syntax.
This both creates and initializes the struct. Each lecture is added directly into the lectures slice.
Beginning at line 23, a for loop is used to iterate over the lectures slice. Dot notation is then used to access the individual fields of the lecture struct.
This is used to read back out the name, the instructor, and the duration of each lecture.
Dot notation is the standard syntax for navigating within structs.

In summary, you have just observed how to create a custom struct collecting related fields into a new data type abstraction,
how you can use composition to create structs that contain other structs,
and how to use dot notation to navigate and access the individual fields within a struct.
*/
