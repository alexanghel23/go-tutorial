package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter float:")
	reader := bufio.NewReader(os.Stdin)
	strNum, err := reader.ReadString('\n')

	if err == nil {
		num, _ := strconv.ParseFloat(strings.TrimSuffix(strNum, "\n"), 64)
		fmt.Println(int(num))

	} else {
		log.Fatal(err)
	}
}
