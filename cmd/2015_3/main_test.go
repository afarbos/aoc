package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/utils"
)

var (
	testDataCountAtLeastOnePresent = map[string]int{
		">":          2,
		"^>v<":       4,
		"^v^v^v^v^v": 2,
	}
	testDataCountAtLeastOnePresentWithRoboSanta = map[string]int{
		"^v":         3,
		"^>v<":       3,
		"^v^v^v^v^v": 11,
	}
)

func TestCountAtLeastOnePresent(t *testing.T) {
	for moves, resExpected := range testDataCountAtLeastOnePresent {
		if housesCount := countAtLeastOnePresent([]byte(moves), false); housesCount != resExpected {
			t.Error("expected ", resExpected, " got ", housesCount)
		}
	}
}

func TestCountAtLeastOnePresentWithRoboSanta(t *testing.T) {
	for moves, resExpected := range testDataCountAtLeastOnePresentWithRoboSanta {
		if housesCount := countAtLeastOnePresent([]byte(moves), true); housesCount != resExpected {
			t.Error("expected ", resExpected, " got ", housesCount)
		}
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)
	for i := 0; i < b.N; i++ {
		main()
	}
}
