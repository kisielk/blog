package main

import (
	"fmt"
	"regexp"
)

var digitsRegexp = regexp.MustCompile(`\d+`)

func main() {
	someString := "1000abcd123"

	// Find just the leftmost match.
	fmt.Println(digitsRegexp.FindString(someString))

	// Find all (-1) the matches.
	fmt.Println(digitsRegexp.FindAllString(someString, -1))
}
