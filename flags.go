// Ben Eggers
// GNU GPL'd

// Demonstrates Go's AMAZING flags library

package main

import (
	"flag"
	"fmt"
)

// The flags can be in any order when this program is run from the command line,
// because Go is cool like that. Also, try giving it a bogus flag, like -aoeuaoeu.
// Flags can be used by running the program ./progname -flagname=flagval ...
func main() {

	// We can define pointers to things by using the flag.TYPE() functions
	strPtr := flag.String("Name1", "default value", "variable usage")
	intPtr := flag.Int("Name2", 0, "But with a different default parameter (duh)")

	// Or, we can have "flag" put things in variables we already have
	var x int
	flag.IntVar(&x, "Name3", 0 /* default value */, "variable usage")
	// Similar functions exist for types like int64, float64, boolean, etc

	// Now that all the flags are defined, we call Parse() to actually get them all
	flag.Parse()

	fmt.Println("str:", *strPtr)
	fmt.Println("int:", *intPtr)
	fmt.Println("other int:", x)

	// We can also get the other things the user puts in as args
	args := flag.Args()
	fmt.Println(args)
}
