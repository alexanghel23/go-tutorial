package main

import (
	"errors"
	"fmt"
)

func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("division by zero not allowed")
	} else {
		return (num1 / num2), nil
	}
}

func main() {
	if result, err := divide(10, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

/*
Functions can be declared to return multiple values. Multiple return values must be declared inside parentheses.
A common pattern often used is to return two values, the functions normal result value, and an error value.
The error value defaults to nil when the function executes without errors. Otherwise it returns information about any errors that were encountered during execution.

In the example shown here, the divide function declared on line six returns two values, the first of which is the result of dividing num one by num two.
The second is a divide by zero error if encountered. Lines 15 to 19 within the main function show you how to call the divide function
and evaluate if it executed without any errors encountered.
It does so by checking the returned error value and checks to see whether it is nil or not and then reacts accordingly.

In summary, you have observed how to declare a function which returns multiple values and how to implement error handling patterns when using functions.
*/
