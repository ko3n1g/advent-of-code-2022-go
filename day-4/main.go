package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Section struct {
	min int
	max int
}

func newSection(section string) *Section {
	pairs := strings.Split(section, "-")
	firstPair, _ := strconv.Atoi(pairs[0])
	secondPair, _ := strconv.Atoi(pairs[1])
	return &Section{min: firstPair, max: secondPair}
}

func isFullOverlap(pair string, result chan int, strictOverlap bool) {
	sections := strings.Split(pair, ",")
	firstSection := newSection(sections[0])
	secondSection := newSection(sections[1])

	var isOverlap bool
	if strictOverlap {
		leftSideOverlap := (firstSection.min <= secondSection.min) && (secondSection.max <= firstSection.max)
		rightSideOverlap := (secondSection.min <= firstSection.min) && (firstSection.max <= secondSection.max)
		isOverlap = leftSideOverlap || rightSideOverlap
	} else {

		// (A.min-A.max) - (B.min-B.max)
		// A.min <= B.min <= A.max <= B.max
		leftSideIntersect := (secondSection.min <= firstSection.max) && (firstSection.min <= secondSection.max)
		rightSideIntersect := (firstSection.min <= secondSection.max) && (secondSection.min <= firstSection.max)
		isOverlap = leftSideIntersect || rightSideIntersect
	}

	var resultAsInt int
	if isOverlap {
		resultAsInt = 1
	} else {
		resultAsInt = 0
	}

	result <- resultAsInt
}

func main() {
	data := getData("challenge.txt")

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