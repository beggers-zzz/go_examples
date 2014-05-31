// Ben Eggers
// GNU GPL'd

// Concurrency stuff in Go. It's so simple!

package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 20; i++ {
		go sqrt(float64(i))
		// `go function_call(args)` creates a new goroutine
		// which runs the function, then exits
	}

	time.Sleep(1) // change this to 0, see what happens!
	fmt.Println("Done")
}

// Newton's method
func sqrt(x float64) {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	fmt.Printf("Sqrt of %d = %f\n", int(x), z)
}
