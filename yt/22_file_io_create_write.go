package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	// close the file
	defer f.Close()

	// list of prime number integers
	iPrimeArr := []int{2, 3, 5, 7, 11}

	// create string arr from int array
	var sPrimeArr []string

	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}

	// cycle through the strings and write them in the file
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Open the file
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Read from the file and print the lines one by one
	scan1 := bufio.NewScanner(f)

	for scan1.Scan() {
		pl("Prime: ", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal()
	}

}
