package main

import (
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
