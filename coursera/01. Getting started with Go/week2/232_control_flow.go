/* Control structures

1. for loops - iterates while a condition is true


2. switch/case (in golang it automatically breaks, no need for `break` statement)

switch x {
case 1:
	fmt.Printf("case1")
case2:
	fmt:Printf("case2")
default:
	fmt.Printf("nocase")
}

// tagless switch
switch {
case x > 1:
	fmt.Printf("case1")
case x < -1:
	fmt.Printf("case2")
default:
	fmt.Printf("nocase")
}

// break - exits the loop

i := 0
for i < 10 {
	i++
	if i == 5 {break}
	fmt.print(i)
}

// continue - skips the current iteration

i := 0
for i < 10 {
	if i == 5 {continue}
	fmt.Printf("hi")
}

