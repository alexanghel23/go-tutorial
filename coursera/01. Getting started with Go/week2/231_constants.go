/*
Constants - expression whose value is known at compile time
Type is inferred from righthand side (boolean, string, number)

const x = 1.3
const (
	y = 4
	z = "Hi"
)
*/

/*
iota - functions used to generate a set of related but distinct constants; often represents a property which has several distinct possible values (like and enumerated type in other languages)
eg. days of the week, month of the year

type Grades int
const (
	A Grades = iota
	B
	C
	D
	F
)


