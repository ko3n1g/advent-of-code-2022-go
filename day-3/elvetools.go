package main

import (
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

func priorityOfItem(item rune) int {
	var offset int
	if unicode.IsUpper(item) {
		offset = 26
	} else {
		offset = 0
	}
	return int(unicode.ToLower(item)-'a'+1) + offset
}

func findCommonItems(firstBag string, secondBag string) map[rune]int {
	// Common items are a mapping of rune to priority of rune
	items := map[rune]int{}

	for _, item := range firstBag {
		// Check if item was already collected
		if _, exists := items[item]; strings.Contains(secondBag, string(item)) && !exists {
			items[item] = priorityOfItem(item)
		}
	}
	return items
}

func sumPrioritiesOfCommonItems(items map[rune]int) int {
	var sumOfPriorities int = 0
	for _, priority := range items {
		sumOfPriorities += priority
	}
	return sumOfPriorities
}

func computeSumOfAllRucksackPriorities(rucksacks []string) int {
	var sumOfPriorities int
	for _, rucksack := range rucksacks {
		sumOfPriorities += sumPrioritiesOfCommonItems(findCommonItems(rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]))
	}
	return sumOfPriorities
}

func computePrioritiesPerGroup(rucksacks []string) int {
	var commonItems string = rucksacks[0]
	var commonItemsAndPriorities map[rune]int
	for _, rucksack := range rucksacks[1:] {
		commonItemsAndPriorities = findCommonItems(commonItems, rucksack)
		commonItems = ""
		for commonItem := range commonItemsAndPriorities {
			commonItems += string(commonItem)
		}
		if commonItems == "" {
			break
		}
	}

	return sumPrioritiesOfCommonItems(commonItemsAndPriorities)
}
