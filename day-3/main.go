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

func getCommonItems(firstBag string, secondBag string) map[rune]int {
	// Common items are a mapping of rune to priority of rune
	items := map[rune]int{}

	for _, item := range firstBag {
		// Check if item was already collected
		if _, exists := items[item]; strings.Contains(secondBag, string(item)) && !exists {
			var offset int
			if unicode.IsUpper(item) {
				offset = 26
			} else {
				offset = 0
			}

			items[item] = int(unicode.ToLower(item)-'a'+1) + offset
		}
	}
	return items
}

func computePriorities(commonItems map[rune]int) int {
	var sumOfPriorities int = 0
	for _, priority := range commonItems {
		sumOfPriorities += priority
	}
	return sumOfPriorities
}

func computeSumOfAllRucksackPriorities(rucksacks []string) int {
	var sumOfPriorities int
	for _, rucksack := range rucksacks {
		sumOfPriorities += computePriorities(getCommonItems(rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]))
	}
	return sumOfPriorities
}

func computePrioritiesPerGroup(rucksacks []string) int {
	var commonItems string = rucksacks[0]
	var commonItemsAndPriorities map[rune]int
	for _, rucksack := range rucksacks[1:] {
		commonItemsAndPriorities = getCommonItems(commonItems, rucksack)
		commonItems = ""
		for commonItem := range commonItemsAndPriorities {
			commonItems += string(commonItem)
		}
		if commonItems == "" {
			break
		}
	}

	return computePriorities(commonItemsAndPriorities)
}

func main() {

	data := getData("challenge.txt")

	// Solve Part 1
	var sumOfPriorities int = computeSumOfAllRucksackPriorities(data)
	fmt.Printf("Sum of priorities is %d\n", sumOfPriorities)

	// Solve Part 2
	var sumOfPrioritiesByGroup int
	const groupSize int = 3
	for i := 0; i < len(data)/3; i++ {

		var sumOfGroup int = computePrioritiesPerGroup(data[groupSize*i : groupSize*(i+1)])
		sumOfPrioritiesByGroup += sumOfGroup
		fmt.Printf("Sum of group %d is %d\n", i+1, sumOfGroup)
	}
	fmt.Printf("Sum of both group's priorities is %d\n", sumOfPrioritiesByGroup)

}
