package main

import "fmt"

func sum(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func minus(num1, num2 int) (result int) {
	result = num1 - num2
	return
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(minus(10, 2))

	result := func(num1 int, num2 int) int {
		sum := sum(num1, num2)
		minus := minus(num1, num2)
		return sum * minus
	}(5, 3)

	fmt.Println(result)
}

/*
Go functions are declared using the func keyword. You may have noticed the main function in all of the previous coding examples presented this as a good example of declaring a function.

Declaring a function allows you to create a reusable block of code, which is typically parameterised with one or several input parameters, and may or may not return one or many values. Function input parameters and return types are optional. A function can heavily be declared without any input and output. In fact, the special main function that is used to start an executable application is a perfect example of this type of function.

Functions are most often declared with package scope, but they can be declared and embedded within other functions, and when they are, they have function scope only. Function naming is also important and can contribute to the visibility of the function. If a functions name begins with an uppercase letter, then the function is considered public and can be called from outside of the package. If, instead, a functions name begins with a lowercase letter, then it can only be called from within the current package file. In the example provided here, the sum and minus functions declared at the top and just beneath the input statement can be called from anywhere within the same file only since both their names begin with a lowercase letter.

Both functions take two input parameters. num1 and num2. And are both typed as integers. The sum function returns a single int value, which is the result of adding num1 and num2 together. Likewise, the minus function returns a single named int value result, which is the result of subtracting num2 from num1. Notice here that the enclosed return statement is called by itself. This is possible since we are declaring a named returned variable in the functions signature. Also note on line 10 that since all input parameters are of the same int type, then we only need to type them once. Moving further down the example, an anonymous function is declared on line 19 inside the main function. This demonstrates a couple of interesting things about functions. First up is that functions can be declared within other functions. Next is that a function can be defined without specifying a name. In this case, they are considered anonymous. This example immediately calls the anonymous function with the input values five and three enclosed in parentheses on line 23.

In summary, you have just observed how to declare private functions with package scope, how to create a function signature, including defining input parameters and return types, how to specify named return types, and how to create an anonymous function within another function and call it immediately.
*/
