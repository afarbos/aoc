package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
	"github.com/afarbos/aoc/pkg/utils"
)

var (
	testDataCountLitLights = map[string]int{
		"toggle 0,0 through 999,0": 1000,
	}

	testDataCountDimLitLights = map[string]int{
		"toggle 0,0 through 999,999": 2000000,
	}
)

func TestCountLitLights(t *testing.T) {
	for instruction, resExpected := range testDataCountLitLights {
		test.EqualInt(t, countLitLights([]string{instruction}), resExpected)
	}
}

func TestCountDimLitLights(t *testing.T) {
	for instruction, resExpected := range testDataCountDimLitLights {
		test.EqualInt(t, countDimLitLights([]string{instruction}), resExpected)
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
