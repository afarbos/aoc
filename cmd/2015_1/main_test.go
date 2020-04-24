package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/utils"
)

var (
	testDataFloor = map[string]int{
		"(())":    0,
		"()()":    0,
		"(((":     3,
		"(()(()(": 3,
		"())":     -1,
		"))(":     -1,
		")))":     -3,
		")())())": -3,
	}

	testDataBasement = map[string]int{
		")":     1,
		"()())": 5,
	}
)

func TestFloor(t *testing.T) {
	for parenthesis, resExpected := range testDataFloor {
		if lastFloor := floor([]byte(parenthesis)); lastFloor != resExpected {
			t.Error("last floor expected", resExpected, "got", lastFloor)
		}
	}
}

func TestBasement(t *testing.T) {
	for parenthesis, resExpected := range testDataBasement {
		if lastFloor := basement([]byte(parenthesis)); lastFloor != resExpected {
			t.Error("first basement expected", resExpected, "got", lastFloor)
		}
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
