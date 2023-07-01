package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(listToSum []int) int {
	var sum = 0
	for _, item := range listToSum {
		sum += item
	}
	return sum
}

func getCaloriesPerElve(file []byte) []int {
	var calories []string = strings.Split(string(file), "\n")
	var caloriesPerElve []int

	var sumOfCalories int = 0
	for _, calorie := range calories {

		if calorie == "" {
			caloriesPerElve = append(caloriesPerElve, sumOfCalories)
			sumOfCalories = 0
			continue
		}

		marks, err := strconv.Atoi(calorie)
		if err != nil {
			panic(err)
		}

		sumOfCalories += marks
	}

	return caloriesPerElve
}

func main() {
	var path string = "/home/okoenig/private_code_projects/advent-of-code-2022-go/data/challenge.txt"
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
