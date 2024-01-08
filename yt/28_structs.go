package main

import (
	"fmt"
)

var pl = fmt.Println

// structs allow you to store values with different data types

type customer struct {
	name    string
	address string
	balance float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.balance)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

func main() {

	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 main st"
	tS.balance = 234.56

	getCustInfo(tS)
	newCustAdd(&tS, "123 South street")
	pl("Address: ", tS.address)

	sS := customer{"Sally Smith", "123 main", 0.0}
	pl("Name:", sS.name)

}
