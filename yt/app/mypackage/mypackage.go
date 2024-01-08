package stuff

import (
	"errors"
	"strconv"
	"time"
)

var Name string = "Derek" // uppercase, will be available outside the package

func IntArrToStrArr(intArr []int) []string { // uppercase, will be available outside the package

	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}

// this part is extended with section 33_getter_setter

type Date struct {
	day   int
	month int
	year  int
}

// setter function to change the date but restrict how we can do it
// they have to capitalized to be accessed outside of the package
func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("Incorrect day value")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(month int) error {
	if (month < 1) || (month > 12) {
		return errors.New("Incorrect month value")
	}
	d.month = month
	return nil
}

func (d *Date) SetYear(year int) error {
	if (year < 1875) || (year > time.Now().Year()) {
		return errors.New("Incorrect year value")
	}
	d.year = year
	return nil
}

// getter functions

func (d *Date) Day() int {
	return d.day
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}
