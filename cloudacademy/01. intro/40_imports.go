package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"

	m "math"
)

func main() {
	name := "cloudacademy"

	fmt.Println(strings.ToUpper(name))
	fmt.Println(uuid.New())
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)
	fmt.Println(m.Round(f))
}

/*
Writing Go code typically requires functionality declared from within other packages. To enable the use of exported types
and/or functions declared and contained within other packages, whether they are provided as part of the standard Go library,
or by 3rd party providers, you use the import statement to import them into the current package.
The import keyword is followed by the import path of the package provided in double quotes.

A package's import path is its module path joined with its subdirectory within the module itself. All of the standard library
packages can be imported using just their package name only. There is no need to declare the module path, since the compiler
already knows where to find them based on the local Go tool chain installation and configuration.

The import keyword can be used to import one or multiple packages at a time. When multiple packages are imported, the import keyword
only needs to be specified once, with the packages being listed out within parenthesis. This style helps by making the code more readable.
Lines 3 through to 7 demonstrate this approach. External packages that have been made public, and of which are quite often stored on GitHub,
can be imported using their GitHub repo URL. The Go tooling is clever enough to connect to GitHub and then download the package locally.
This can be seen in the example here on line 6, where I'm utilizing Google's UUID package.

Imported packages can also be aliased to avoid package naming conflicts. In the case that there are two packages, which share the same name
but for which come from different locations and implement different things, an alias can be assigned to one or both to ensure that when their
functionality is called upon, the right code gets compiled and executed. For example, I have aliased the math package on line 10 with the m alias.

Now, if you look closely at all of the functions that are invoked on any of the imported packages, you will see that they all begin
with a capital letter. The functions I'm referring to are ToUpper, New, ParseFloat, and Round, and as just stated they all start with a capital letter.
It is this convention that actually makes them exported, meaning that they can be seen from outside of the package that they were declared within
and can be imported into and invoked from within other packages. This will be covered in more detail as we progress.

In summary, you've observed how to use the import keyword to import packages, how to group multiple import statements together,
how to use aliases in import statements, how aliases can be used to alias imported packages, and how to import external public packages hosted on GitHub.
*/
