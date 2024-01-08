package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {

	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")

	sV2 := replacer.Replace(sV1)

	pl("Length:", len(sV2))

	pl("Contains another?", strings.Contains(sV2, "word"))

	pl("o index", strings.Index(sV2, "o"))

	pl("Replace", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\nSome Words\n"
	pl(sV3)
	sV3 = strings.TrimSpace(sV3)

	pl(sV3)

	pl("Split", strings.Split("a-b-c", "-"))

	pl("Lower", strings.ToLower(sV2))
	pl("Lower", strings.ToUpper(sV2))

	pl("Prefix", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix", strings.HasSuffix("tacocat", "cat"))

}
