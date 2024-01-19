// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	var name string
// 	var address string
// 	fmt.Print("Enter name: ")
// 	fmt.Scan(&name)
// 	fmt.Print("Enter address: ")
// 	fmt.Scan(&address)
// 	fmt.Println()

// 	person := map[string]string{
// 		"name":    name,
// 		"address": address,
// 	}

// 	out, err := json.Marshal(person)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(string(out))
// 	}
// }

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter the name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Please enter the address: ")
	addr, _ := reader.ReadString('\n')
	addr = strings.TrimSpace(addr)
	m := map[string]string{
		"name":    name,
		"address": addr,
	}
	person, err := json.Marshal(m)
	if err == nil {
		fmt.Printf("Printed as bytes: %v\n", person)
		fmt.Printf("Printed as string: %s", string(person))
	} else {
		fmt.Printf("Invalid input!")
	}
}
