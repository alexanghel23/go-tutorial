package main

import (
	"fmt"
	"os"
)

/*
Files
- Linear access, not random access -> mechanical delay (just like a video tape)
- Basic operations:
	- Open - get handle for access
	- Read - read bytes into []byte
	- Write - write []byte into file
	- Close - release handle
	- Seek - move read/write head

ioutil File read
- io/ioutil package has basic functions (deprecated)
- os.ReadFile is the new function
- explicit open/close are not needed
- large files cause a problem so os.ReadFile might not be appropriate

*/

func main() {

	dat, e := os.ReadFile("test.txt") //dat is []byte filled with contents of the entire file

	if e == nil {
		fmt.Println(dat)
	}

	var dat2 string = "Hello world"
	err := os.WriteFile("outfile.txt", []byte(dat2), 0777)
	if err == nil {
		fmt.Println("File created")
	}

}
