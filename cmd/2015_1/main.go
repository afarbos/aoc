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
	if lastFloor := floor(parenthesis); lastFloor != resFloor {
		log.Fatal("Expected last floor ", resFloor, " got ", lastFloor)
	} else {
		log.Println(lastFloor)
	}
	if firstBasement := basement(parenthesis); firstBasement != resBasement {
		log.Fatal("Expected first basement ", resBasement, " got ", firstBasement)
	} else {
		log.Println(firstBasement)
	}
}
