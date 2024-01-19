package main

import (
	"fmt"
	"os"
)

/*

os package file access
- os package provides more control than ioutil
- functions:
	- os.Open() - returns a file descriptor
	- os.Close() - closes a file
	- os.Read() - reads from a file into a []byte; you can control the amount of read in comparison with ioutil
	- os.Write()
*/

func main() {

	// Read
	f, _ := os.Open("test.txt")
	barr := make([]byte, 10)

	nb, _ := f.Read(barr) // read first 10 bytes

	fmt.Println(nb)

	f.Close()

	// Write
	f2, _ := os.Create("outfile2.txt")
	barr2 := []byte{1, 2, 3}
	f2.Write(barr2)
	f2.WriteString("Hi")

}
