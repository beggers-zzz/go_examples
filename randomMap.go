// Ben Eggers
// GNU GPL'd

// Demonstrates random map iteration in Go (proving it to myself)

package main

import "fmt"

func main() {

	things := map[rune]int{} // our demo map--'rune' is go's term for unicode character
	numChars := 10           // number of characters to print
	initChar := 60           // character to start on

	// populate
	for i := initChar; i < initChar+numChars; i++ {
		things[rune(i)] = i
	}

	// and print. Interestingly, the order seems to be biased towards sorted
	// order
	for character, charCount := range things {
		fmt.Println(string(character), ":", charCount)
	}
}
