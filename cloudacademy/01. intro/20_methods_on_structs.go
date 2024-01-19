package main

import "fmt"

type person struct {
	firstname string
	surname   string
	age       int
}

func (p *person) fullname() string {
	return p.firstname + " " + p.surname
}

func (p *person) canDrive() bool {
	if p.age >= 20 {
		return true
	} else {
		return false
	}
}

func (p *person) updateAge(newAge int) {
	p.age = newAge
}

func main() {
	person1 := person{"John", "Doe", 40}
	person2 := person{"Mark", "Doe", 19}

	fmt.Printf("%s can drive: %t\n", person1.fullname(), person1.canDrive())
	fmt.Printf("%s can drive: %t\n", person2.fullname(), person2.canDrive())

	person2.updateAge(person2.age + 1) //marks birthday
	fmt.Println(person2.age)

	fmt.Printf("%s can drive: %t\n", person2.fullname(), person2.canDrive())
}

/*
Although Go isn't strictly an object oriented program language like Java and C# are, it still provides various language features that allow for an object oriented style of programming.

One such feature is the concept of a method.
A Go method is defined and associated with a receiver type allowing you to use dot notation to invoke the method on the receiver type.
Let's see how this works.
For starters, methods in Go are identical to functions syntaxwise with the one exception of an added receiver argument.
The receiver in most cases will be declared as a struct type.
Having said that, receivers can also be non-struct types as you'll see in the next part.
The receiver argument is declared with a name and type, inserts inside parentheses appearing immediately after the func keyword and before the method name.

A method receiver argument can be declared as a pointer or non-pointer type.
There are two primary reasons to use a pointer receiver.
The first is so that the method can modify the data that its receiver points to and
the second is to avoid copying the value on each method core which as we know is a more efficient approach if the receiver is a large struct type.
In the example provided here, all of them point to three methods:
The fullname method, implemented on lines 11 to 13, the canDrive method, implemented on lines 15 to 21, and the updateAge method, implemented on lines 23, 24 and 25.

All methods have as the receiver a variable named p which is a pointer to a person type.
The person data type as declared earlier as a struct spanning on lines five through to nine.
The fullname method simply concatenates the person's first name and surname and returns it as the person's full name.
The canDrive method is used to test the person's age to see if they are 20 years or older and if so, then this method returns true.
Otherwise it returns false. Methods are invoked using dot notation on the variable as long as the variable's datatype matches the method's receiver datatype.

This can be seen on line 31 for person one and again on line 32, 34 and 37 for person two.
This style of coding has the benefit of being easier to read and navigate.
The receiver argument can be either a pointer to a type or just a plain type.
With the pointer the a type option, having the advantage of allowing the method to modify the past and data such that the modifications last beyond the execution of the method.
Knowing that the fullname in candDrive methods perform read-only operations on the underlined person struct, they then could in fact be declared to use non pointer receiver types like so.

In summary, you've observed methods are identical to functions, except for the added receiver argument inserted between the func keyword and the method's name.
Receiver arguments are often custom struct types. Receiver arguments can be either value or reference types and dot notation is the preferred way to invoke a method.
*/
