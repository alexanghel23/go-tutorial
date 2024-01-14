package main

import "fmt"

func main() {
	var num1 int = 100
	var num2 int = 200
	var str1 string = "blah"

	var ptr1 *int = &num1
	//var ptr1 *int = &str1 //compile error

	fmt.Printf("mem address of num1 is %p\n", &num1)
	fmt.Printf("mem address of num2 is %p\n", &num2)
	fmt.Printf("mem address of str1 is %p\n", &str1)

	fmt.Printf("ptr1 points to mem address %p\n", ptr1)

	*ptr1 = 101
	fmt.Println(num1)

	ptr1 = &num2
	fmt.Printf("ptr1 points to mem address %p\n", ptr1)
	fmt.Println(*ptr1)

	ptr2 := new(int)
	ptr2 = ptr1
	fmt.Println(*ptr2)
}

/*
Pointers and Go provide you with the ability to store the memory address of another variable.
Why would you want to do this you may be asking yourself.
I'll answer this in the next pointer based demo.
But, for now let me continue explaining what they are and how you code with them.
For starters, any variable you declare in a Go program requires memory to store the data or value assigned to the variable.
The memory requirements in terms of size and space are entirely dependent on the type of the variable.
Memory within a computer system is addressable and variables are used to meet themselves to the addressable location where the assigned data resides.

A pointer is also a variable, but a special kind of one. That is because it stores a memory address rather than the data itself.
The memory address is often the address of another variable. Let's quickly run this example and then next explain its inner workings.
In the example provided here, I declare num one and num two as standard integer typed variables on lines six and seven and I assigned the values 100 and 200 respectively.
Line eight declares the string one variable, which is initialized with the string value "blah."
A pointer to an integer is declared on line 10 and is initialized with the memory address of the num one variable.

Pointers are declared by prefixing an asterisks to the type that the pointer is pointing to.
An ampersand is prefixed to a variable to get its memory address.
It is important to realize that in this case, pointer one can only store the memory address of int variables.
If you were to try and assign the memory address of a different data type, then a compiled time error would be thrown.
Lines 13, 14 and 15 print out the memory addresses for the num one, num two, and string one variables, respectively.
Line 17 then prints out the memory address stored in pointer one. Notice that it stores the same num one memory address.
We can then de=reference pointer one by prefixing it with an asterisk, as is done on line 19, and change its stored value.
In this case, we are changing the value from 100 to 101.

Now the cool thing about doing this is that any other variable that points to the same memory location will see this update.
Line 20 prints out the num one variable and as expected, we can see that it prints out 101.
Pointers can be updated to point at different memory addresses, even after initialization, as shown on line 22.
Here, pointer one now points at the memory address of the num two variable.
When we de-reference pointer one again, this time it prints out 200, which is the value held in the memory location of variable num two.
Pointers can also be created using the in built new function, as shown on line 26.
Here, pointer two is created and initialized using shorthand notation, and is declared to also be a pointer to an int.
Line 27 simply tells pointer two to point to the same memory address that pointer one currently points to.

Finally, if we were to de-reference pointer two, it will again print out 200, which is the value held in the memory location of variable num two.

In summary, you have just observed that pointers are variables that store memory addresses of other variables.
Pointers are typed.
A memory address for a given variable is found by prefixing the variable's name with an ampersand, and how to create a new pointer using the built in new function.
The new function speaks a type as an input parameter, and then allocates just enough memory to values of that type.
A couple of notes about pointers.
The zero value of a pointer is null.
This means any uninitialized pointer will have the value null, and there is no pointer arithmetic and go.
*/
