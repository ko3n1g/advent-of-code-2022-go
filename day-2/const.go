package main

var encoder = map[string]Figure{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var beats = map[Figure]Figure{
	Rock:     Scissors,
	Scissors: Paper,
	Paper:    Rock,
}
