package main

import (
	"errors"
	"fmt"
	"math"
)

func circleArea(radius float32) (float32, error) {
	if radius < 0 {
		return 0, errors.New("radius should be positive value")
	} else {
		return math.Pi * radius * radius, nil
	}
}

func main() {
	area1, _ := circleArea(3)
	fmt.Println(area1)

	if area2, err := circleArea(-3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area2)
	}
}

/*
When building large and complex Go programs it's always difficult to anticipate all of the possible outcomes that can occur when the program runs.
Error conditions exist and should be handled gracefully. Go provides an ability to raise and communicate errors through the built in error
interface type which is managed via several useful functions found within the errors package.

The mechanism in which errors are communicated back up the call chain happens to be via a functions capability of returning multiple values,
one of which will be the error itself. Back within the callee scope - error conditions can then be checked upon and actions taken accordingly
using the same language features already reviewed. In most cases checks will be done using a simple if/else statement that tests to see if the
returned error variable is nil or not. In the example provided here - we begin by importing the "errors" package on line 4. Next the circleArea
function is declared on lines 7 to 13 - and for which performs a simple validation check on the incoming radius parameter to ensure that its value is positive.

The function signature has a multi return value - composed inside the last set of parentheses and consisting of the area float32 and an error which is of type error.
In the case that the radius parameter is not a positive value - an error is created and returned back to the caller - this can be seen on line 9.
In the case that the functions execution runs without errors, meaning the radius is a positive number then the error return value is set to nil as can be seen on line 11.
Within the main function on line 16, the circleArea function is called with the radius set to 3.
In this case since I have knowledge that the circleArea function will execute without causing any errors.

I can simply ignore the returned error value by assigning it to an underscore. Lines 19 to 23 however, demonstrates an invocation of the circleArea function
with a negative radius - in this case -3. The invocation of the circleArea function is done inline within an if/else statement.
An explicit test is done against the returned error variable to test whether it contains an error or not. Not being the case when it is 6 nil.
Let's run this example and examine the resulting output. Here we can see that the first call to the circleArea function succeeded
with the result 28.27 and that the second call to the circleArea function failed with the error message "radius should be positive" message.

In summary, you've just observed How to import the errors package to give you the ability to create errors How to return errors
using the errors .New function And how to test for returned errors and react accordingly
*/
