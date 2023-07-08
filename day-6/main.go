package main

import (
	"fmt"

	"github.com/ko3n1g/advent-of-code-2022-go/day-6/marker"
	"github.com/ko3n1g/advent-of-code-2022-go/tools"
)

func main() {
	data := tools.GetData("example.txt")
	const packetSize int = 4
	const messageSize int = 14

	for i, row := range data {
		var idx int
		var err error

		idx, err = marker.GetIdxOfMarker(row, packetSize)
		if err != nil {
			fmt.Printf("%s at message %d", err, i+1)
			continue
		}
		fmt.Printf("start-of-packet at idx %d\n", idx)

		idx, err = marker.GetIdxOfMarker(row, messageSize)
		if err != nil {
			fmt.Printf("%s at message %d", err, i+1)
			continue
		}
		fmt.Printf("start-of-message at idx %d\n\n", idx)

	}
}
