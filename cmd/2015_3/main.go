package main

import (
	"flag"
	"image"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	resHousesCount     = 2565
	resHousesCountRobo = 2639
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func countAtLeastOnePresent(moves []byte, roboSanta bool) int {
	houses := make(map[image.Point]struct{})
	currentPosition, roboSantaPosition := image.Point{0, 0}, image.Point{0, 0}
	houses[currentPosition] = struct{}{}

	for i, move := range moves {
		pt := &currentPosition
		if roboSanta && i%2 == 1 {
			pt = &roboSantaPosition
		}
		switch move {
		case "^"[0]:
			pt.X++
		case "v"[0]:
			pt.X--
		case "<"[0]:
			pt.Y--
		case ">"[0]:
			pt.Y++

		}
		houses[*pt] = struct{}{}
	}
	return len(houses)
}

func main() {
	flag.Parse()
	moves := read.Bytes(flagInput)
	utils.AssertEqual(countAtLeastOnePresent(moves, false), resHousesCount)
	utils.AssertEqual(countAtLeastOnePresent(moves, true), resHousesCountRobo)
}
