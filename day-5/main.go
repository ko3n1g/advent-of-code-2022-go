package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	board "github.com/ko3n1g/advent-of-code-2022-go/day-5/board"
	"github.com/ko3n1g/advent-of-code-2022-go/day-5/slots"
)

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func playGame(myBoard board.Board, instructions []string) {
	for _, move := range instructions {
		fmt.Println(move)
		var (
			quantity int
			from_idx int
			to_idx   int
		)
		_, err := fmt.Fscanf(strings.NewReader(move), "move %d from %d to %d", &quantity, &from_idx, &to_idx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
		}
		for i := 0; i < quantity; i++ {
			myBoard.MoveFigure(from_idx-1, to_idx-1)
		}
		fmt.Println(myBoard)
		fmt.Println()
	}
}

func main() {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(currPath + "/data/" + "challenge.txt")
	if err != nil {
		panic(err)
	}

	game := strings.Split(string(file), "\n\n")

	boardDescription, instructions := strings.Split(game[0], "\n"), strings.Split(game[1], "\n")
	ReverseSlice(boardDescription)

	const numCols = 9
	var myBoard board.Board = make([]slots.Slot, numCols)
	myBoard.ParseFromList(boardDescription, numCols)

	playGame(myBoard, instructions)

	fmt.Printf("\nFinal solution:\n")
	var solution string = ""
	for i := 0; i < numCols; i++ {
		solution += myBoard[i].Pop()
	}
	fmt.Println(solution)
}
