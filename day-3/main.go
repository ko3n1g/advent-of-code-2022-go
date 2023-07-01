package main

import (
	"fmt"
)

func main() {

	data := getData("challenge.txt")

	// Solve Part 1
	var sumOfPriorities int = computeSumOfAllRucksackPriorities(data)
	fmt.Printf("Sum of priorities is %d\n", sumOfPriorities)

	// Solve Part 2
	var sumOfPrioritiesByGroup int

	for i := 0; i < len(data)/3; i++ {

		var sumOfGroup int = computePrioritiesPerGroup(data[groupSize*i : groupSize*(i+1)])
		sumOfPrioritiesByGroup += sumOfGroup
		fmt.Printf("Sum of group %d is %d\n", i+1, sumOfGroup)
	}
	fmt.Printf("Sum of both group's priorities is %d\n", sumOfPrioritiesByGroup)

}
