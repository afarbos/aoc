package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
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
		test.EqualInt(t, floor([]byte(parenthesis)), resExpected)
	}
}

func TestBasement(t *testing.T) {
	for parenthesis, resExpected := range testDataBasement {
		test.EqualInt(t, basement([]byte(parenthesis)), resExpected)
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
