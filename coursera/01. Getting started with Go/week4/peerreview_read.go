// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Person struct {
// 	fname string
// 	lname string
// }

// func main() {
// 	var out []Person
// 	var filename string

// 	fmt.Print("Enter file name: ")
// 	fmt.Scanf("%s", &filename)
// 	fmt.Println("File name provided is", filename)

// 	// filename = "test.txt"

// 	file, _ := os.Open(filename)

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		ls := strings.Split(scanner.Text(), " ")
// 		out = append(out, Person{fname: ls[0], lname: ls[1]})
// 	}
// 	file.Close()

// 	for _, p := range out {
// 		fmt.Print(p.fname, " ", p.lname, "\n")
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	personlist := make([]Person, 0, 3)
	var fileName string
	fmt.Println("\n Enter the file name present in current directory")
	_, error := fmt.Scanf("%s", &fileName)
	check(error)
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var aName Person
		aName.fname, aName.lname = s[0], s[1]
		personlist = append(personlist, aName)
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range personlist {
		fmt.Println(eachline.fname, eachline.lname)
	}
}
