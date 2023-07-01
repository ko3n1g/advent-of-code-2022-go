package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getData(fileName string) []string {
	currPath, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(currPath + "/data/" + fileName)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(file), "\n")
}

func computePriorities(rucksacks []string) int {
	var sumOfPriorities int
	for _, rucksack := range rucksacks {

		// Get first and second compartment, each are of same size
		firstCompartment, secondCompartment := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]

		// Common items are a mapping of rune to priority of rune
		items := map[rune]int{}

		for _, item := range firstCompartment {

			// Check if item was already collected
			if _, exists := items[item]; strings.Contains(secondCompartment, string(item)) && !exists {
				var offset int
				if unicode.IsUpper(item) {
					offset = 26
				} else {
					offset = 0
				}

				items[item] = int(unicode.ToLower(item)-'a'+1) + offset
				sumOfPriorities += items[item]
			}
		}

	}
	return sumOfPriorities
}

func main() {

	data := getData("challenge.txt")

	var sumOfPriorities int = computePriorities(data)

	fmt.Printf("Sum of priorities is %d\n", sumOfPriorities)

}
