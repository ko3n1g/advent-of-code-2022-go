package main

import (
	"fmt"
	"os"
	"strings"
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

func computeScore(game []string) int {
	var score int
	for _, round := range game {

		if round == "" {
			continue
		}

		figures := strings.Split(round, " ")
		opponent, ours := encoder[figures[0]], encoder[figures[1]]

		var subScore int
		if beats[opponent] == ours {
			subScore = 0
		} else if opponent == ours {
			subScore = 3
		} else {
			subScore = 6
		}

		score += int(ours) + subScore
	}
	return score
}

func main() {

	data := getData("challenge.txt")

	// Compute score
	var score int = computeScore(data)

	// Print our score
	fmt.Printf("Score of all rounds is: %d\n", score)
}
