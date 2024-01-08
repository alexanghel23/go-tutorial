package main

import (
	"fmt"
	"time"
)

const (
	YYYYMMDD = "2006-01-02"
)

var pl = fmt.Println

func main() {

	now := time.Now()

	pl(now.Year(), now.Month(), now.Day())

	pl(now.Hour(), now.Minute(), now.Second())

	pl(now.Format(YYYYMMDD))

}
