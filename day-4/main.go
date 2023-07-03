package main

import (
	"fmt"
	"sync"

	"github.com/ko3n1g/advent-of-code-2022-go/tools"
)

func main() {
	var wg sync.WaitGroup
	data := tools.GetData("challenge.txt")

	sumOfOverlaps := make(chan int, 1)
	sumOfOverlaps <- 0
	for idx, pair := range data {
		wg.Add(1)
		go isFullOverlap(idx, pair, sumOfOverlaps, true, &wg)
	}
	wg.Wait()
	fmt.Printf("Overlaps %d\n", <-sumOfOverlaps)

	sumOfOverlapsNonStrict := make(chan int, 1)
	sumOfOverlapsNonStrict <- 0
	for idx, pair := range data {
		wg.Add(1)
		go isFullOverlap(idx, pair, sumOfOverlapsNonStrict, false, &wg)
	}
	wg.Wait()
	fmt.Printf("Overlaps %d\n", <-sumOfOverlapsNonStrict)
}
