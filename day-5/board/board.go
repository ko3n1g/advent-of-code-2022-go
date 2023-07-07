package board

import (
	"unicode"

	"github.com/ko3n1g/advent-of-code-2022-go/day-5/slots"
)

type Board []slots.Slot

func (board *Board) MoveFigure(from_idx int, to_idx int, nItems int) {
	(*board)[to_idx].Push((*board)[from_idx].Pop(nItems))
}

func (board *Board) ParseFromList(listOfBoardDescriptions []string, nCols int) {
	for i := 0; i < nCols; i++ {
		for _, row := range listOfBoardDescriptions {
			item := rune(row[i*4+1])
			if !unicode.IsLetter(item) {
				continue
			}

			(*board)[i].Push([]string{string(item)})
		}
	}
}
