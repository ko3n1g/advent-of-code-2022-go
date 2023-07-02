package main

import (
	"fmt"

	"github.com/ko3n1g/advent-of-code-2022-go/tools"
)

func main() {
	data := tools.GetData("challenge.txt")

	// This time, we're gonna use a go-routine for multiprocessing!
	// For each sample, get the min/max values of each pair's section
	// Fully overlap's simple then
	// How about a struct as a container?

	// Solution of Part 1
	sumOfOverlaps := 0
	for _, pair := range data {
		result := make(chan int)
		go isFullOverlap(pair, result, true)
		sumOfOverlaps += <-result
	}

	fmt.Printf("Overlaps %d\n", sumOfOverlaps)

	// I very much believe that the overhead of creating threads (as well as
	// communicating with them) will not be rewarded in this use case.
	// However, it is a nice opportunity to familiarize oneself with
	// Go-routines.

	// Solution of Part 2
	sumOfOverlapsNonStrict := 0
	for _, pair := range data {
		result := make(chan int)
		go isFullOverlap(pair, result, false)
		sumOfOverlapsNonStrict += <-result
	}

	fmt.Printf("Overlaps %d\n", sumOfOverlapsNonStrict)

}
