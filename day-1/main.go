package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var path string = "/home/okoenig/private_code_projects/advent-of-code-2022-go/day-1/data/challenge.txt"
	file, err := os.ReadFile(path)

	// Panic if reading file failed
	if err != nil {
		panic(err)
	}

	// Get calories per elve
	caloriesPerElve := getCaloriesPerElve(file)

	// Sort ascending
	sort.Ints(caloriesPerElve)

	// Part 1: Print calories of Elve carrying the most
	fmt.Printf("Elve carrying the most calories got %d in da bag.\n", caloriesPerElve[len(caloriesPerElve)-1])
	// Part 2: Print calories of Top-3 Elves
	fmt.Printf("Top 3 elves are carrying %d calories.\n", sum(caloriesPerElve[len(caloriesPerElve)-3:]))
}
