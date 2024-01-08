package main

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {

	// reStr := "The ape was at the apex"

	// match, _ := regexp.MatchString("(ape[^ ]?", reStr) // looking for ape not follwed by a space
	// pl(match)

	reStr := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at")
	pl("Matchstring: ", r.MatchString(reStr))
	pl("FindString: ", r.FindString(reStr))
	pl("Index: ", r.FindStringIndex(reStr))
	pl("All Strings", r.FindAllString(reStr, -1))
	pl("First two strings", r.FindAllString(reStr, 2))
	pl("All Submatch Index", r.FindAllStringSubmatchIndex(reStr, -1))
	pl(r.ReplaceAllString(reStr, "Dog"))
}
