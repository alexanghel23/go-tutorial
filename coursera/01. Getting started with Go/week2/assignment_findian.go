package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var text string

	fmt.Println("Enter string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	text = strings.ToLower(text)

	hasPreffix := strings.HasPrefix(text, "i")
	hasSuffix := strings.HasSuffix(text, "n")
	contains := strings.Contains(text, "a")
	if hasPreffix && hasSuffix && contains == true {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

	// more complicated solution

	// reader := bufio.NewReader(os.Stdin)
	// strVal, err := reader.ReadString('\n')

	// if err == nil {
	// 	var arr = []string{}

	// 	for _, c := range strings.ToLower(strings.TrimSuffix(strVal, "\n")) {

	// 		arr = append(arr, string(c))
	// 		// fmt.Printf("%d %c\n", i, c)

	// 	}

	// 	var checkA bool = false
	// 	for _, c := range arr[1 : len(arr)-1] {
	// 		if c == "a" {
	// 			checkA = true
	// 			break
	// 		}
	// 	}

	// 	if arr[0] == "i" && arr[len(arr)-1] == "n" && checkA {
	// 		fmt.Println("Found!")
	// 	} else {
	// 		fmt.Println("Not Found!")
	// 	}

	// } else {
	// 	log.Fatal(err)
	// }
}
