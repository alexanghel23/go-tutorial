package main

import "fmt"

func notString(msg string) {
	msg = fmt.Sprintf("not%s", msg)
}

func notStringPtr(msg *string) {
	*msg = fmt.Sprintf("not%s", *msg)
}

func main() {
	message := "cloudacademy"

	notString(message)
	fmt.Println(message)

	notStringPtr(&message)
	fmt.Println(message)
}

/*
 In the previous demonstration, I showed you the basics of working with pointers, how to create and initialize them, how to de reference them, and how to re point them, etc.
 In this demonstration, I want to explain why they are useful and the situations that you should consider using them.
 In the example provided here, I've implemented two string based functions, notString and notStringPtr.
 Both functions simply take a single input parameter of type string.
 The function is then designed to simply attach the string "not" to the front of the input string.

Note here that both functions do exactly that, and nothing else, they don't even have a return value.
In the main function a message variable is created and initialized with the string "cloudacademy".
We then call the first function, the notString function on line 16, passing in the message variable.
When execution returns from this function, we immediately print out the message variable to find that it still holds the same initial string "cloudacademy".
This function and in particular its' input parameters use a copy by value approach, that is the parameters as used within the function are copies of the variables that are passed into the function.
Next, we call the second function, the notStringPtr function on line 19, again passing in the same "cloudacademy" message.
When execution returns from this function, we immediately print out the message variable to find that it now contains the updated string "notcloudacademy".

So why the change in behavior?
Well, if you take a closer look at the 2 different function implementations, in particular the signatures, you will find that the notStringPtr function actually takes in a pointer to a string input parameter.
The message input parameter is then de referenced with the string concatenation operation now taking place on the exact same string held in the exact same memory location that the original message variable declared in the main function points at.
This function and in particular its' input parameters use a copy by reference approach, that is the parameters as used within the function are direct references to the variables that are passed into the function.

Declaring an input parameter as a pointer to a data type, allows you to manipulate the contents of that parameter within the function and then to have those updates available back outside of the function.
This is one of the main motives for using pointers. Another area where pointers are beneficial, is the scenario where you have a variable that holds a very large amount of data and you need to pass this variable into a function.
Using a pointer prevents you from having to recopy and duplicate all of the data, keeping your memory space free.

In summary, you've just observed that input parameters can be typed as pointers, and that this allows the function to modify the data stored within the parameters such that the modification are then available outside of the function.
*/
