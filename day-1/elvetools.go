package main

import (
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
