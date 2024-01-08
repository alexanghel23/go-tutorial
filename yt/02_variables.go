package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {

	// var name type

	var vName string = "Alex"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"

	pl(vName, v1, v2, v3)

	pl(reflect.TypeOf("Alex"))

	cV1 := 1.5
	cV2 := int(cV1)

	pl(cV1)
	pl(cV2)

	cV1 = 2 // cannot reasign new values using :=

	pl(cV1)

	cV3 := "500000000"

	cV4, err := strconv.Atoi(cV3) // convert string to int (Atoi = AsciiToInteger)
	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := 500000000

	cV6 := strconv.Itoa(cV5) // convert int to string (Itoa = IntegerToAscii)
	pl(cV6, reflect.TypeOf(cV6))

	cV7 := "3.14"

	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

}
