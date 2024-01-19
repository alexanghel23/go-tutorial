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

//unexported
func getX() int {
	return x
}

-- main.go --
package main

import (
	"github.com/cloudacademy/go/mod/demo/util"
	"fmt"
)

func main() {
	fmt.Println(util.X)
	//fmt.Println(util.x) //compile failure - unexported

	fmt.Println(util.GetX())
	//fmt.Println(util.getX()) //compile failure - unexported

}

/*
As previously mentioned, the Go language specification uses a naming convention whereby variables, constants, 
functions, et cetera, that you want to export and make available to be imported elsewhere are named with a starting capital letter. 

Anything that begins with a lowercase letter is determined to be unexported and therefore not importable. In other 
languages such as C# and Java, this type of control is implemented using dedicated modified keywords such as public and private. 
To see this in action, consider the example provided here. For starters, this example is a multi-file example, the first of which 
we've seen and used so far. The double-dash lines are the demarkation points for the start and end of each file. These are located on lines one, four, and 20. 
You will only require these to model and test multi-file designs within the Go Playground. They are not considered part of the formal Go language syntax. 

In this example we have three files. The files go.mod and main.go are both located in the root directory of the example while the x.go file is 
located in a directory named util. The go.mod file is used to declare a module which is used to organize and contain the util package which 
is later imported within the main.go file. The x.go utility file is declared to be in the util package as per the package statement at the top. 
Within this file, there are two constants declared on lines seven and eight, one named with an uppercase X, ensuring that it is an exported constant, 
and the other with a lowercase x which makes it an unexported constant. 

Next up are two functions declared on lines 11 to 13 and lines 15 to 18. Both functions share the same name, albeit the first beginning with a 
capital G making it exported and the other one starting with a lowercase g making it unexported. The main.go file which begins on line 21 imports 
the util package contained in the module declared above. Within the main function, the util.X constant is accessed on line 29. 
This is possible since it was made exportable within the util package as earlier explained. The equivalent lowercase x constant if accessed would not 
be possible and would actually cause a compile time error if attempted. Likewise the same behavior happens for the invocation of the GetX functions on lines 32 and 33. 

The point here is that the naming convention dictates whether it is exported or not and therefore whether it can be visible in other packages. 
Running this example returns the values 100 and 200. Let's alter the program by attempting to access the unexported constant and function like so. 
Rerunning the program now throws a compile time error as expected. Note that this is a compile time error, not a run time error. 

In summary, you have observed that the starting letter for variables, constants, functions, et cetera, dictates whether they are exported or not 
therefore whether they are visible in other packages. If the name starts with a capital letter, then it is exported otherwise not.
*/