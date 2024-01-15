-- go.mod --
module github.com/cloudacademy/go/mod/demo

-- util/x.go --
package util

const X int = 100 //exported
const x int = 200 //unexported

//exported
func GetX() int {
	return x
}

//exported
func GetXY() int {
	return x * getY() //visible due to being in same package
}

-- util/y.go --
package util

const y int = 300 //unexported

//unexported
func getY() int {
	return y
}

-- main.go --
package main

import (
	"github.com/cloudacademy/go/mod/demo/util"
	"fmt"
)

func main() {
	fmt.Println(util.GetXY())
	
	//fmt.Println(util.getY()) //compile failure - unexported
}

/*
When it comes to understanding the visibility of variables, constants, and/or functions et cetera, it's important to understand 
that any of them declared outside of any function are visible across the entire package in which they were declared, regardless of 
whether their names are capitalized or not. 

To put it another way, all files that belong to the same package have visibility of them. This is irrespective of the exporting naming 
convention previously reviewed. Therefore, any variable, constant, and/or function whose name starts with either a lowercase or uppercase 
character will be visible and available in all other source files that share the same starting package name statement. In the example provided here, 
we can see that the util package which now consists of two source files, the x.go and y.go files. File x.go spans lines five to 19, and file y.go spans lines 21 to 29. 

Both of these files belong to the same util package as per the package statements at the top of each file. The y.go file declares an 
unexported getY function on lines 26 to 28. Even though this function is unexported, it can still be called upon within the x.go file as 
demonstrated within the GetXY function declared on lines 16 to 18. Again, the reason for this is that both the x.go and y.go files belong 
to the same util package. Therefore, each can see all of the each other's variables, constants, and/or functions et cetera, that have been 
declared outside of any other function. Running this example produces the result 60,000. 

The explanation for this follows. The main function makes a call to the exported util.GetXY function on line 39. This function returns the 
result of multiplying little x, a constant whose value is 200, by the unexported but still visible getY function. This is the key point of this example. 
The fact that the getY function is still visible within the x.go file, even though it was declared in the y.go file and by naming convention is unexported. 
The reason that this is all possible is due to the fact that both files, x.go and y.go belong to the same util package. 
The getY function returns the value 300, which is the value stored in the constant y. Finally resulting in the multiplication of 200 times 300, which we know equals 60,000. 

In summary, you've observed source files belonging to the same named package can see each other's variables, constants, and/or functions et cetera, if they were 
defined outside of other functions, and regardless of whether they were exported or not.
*/