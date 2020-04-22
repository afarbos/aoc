package main

import (
	"flag"
	"log"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	resFloor    = 138
	resBasement = 1771
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func floor(b []byte) int {
	floor := 0
	for _, b := range b {
		if b == "("[0] {
			floor++
		} else if b == ")"[0] {
			floor--
		}
	}
	return floor
}

func basement(b []byte) int {
	floor := 0
	for i, b := range b {
		if b == "("[0] {
			floor++
		} else if b == ")"[0] {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	log.Fatal("Did not find the position")
	return 0
}

func main() {
	flag.Parse()
	parenthesis := read.Bytes(flagInput)
	utils.AssertEqual(floor(parenthesis), resFloor)
	utils.AssertEqual(basement(parenthesis), resBasement)
}
