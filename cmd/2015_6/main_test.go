package main

import (
	"testing"

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
		if res := countLitLights([]string{instruction}); res != resExpected {
			t.Error("expected ", resExpected, " got ", res)
		}
	}
}

func TestCountDimLitLights(t *testing.T) {
	for instruction, resExpected := range testDataCountDimLitLights {
		if res := countDimLitLights([]string{instruction}); res != resExpected {
			t.Error("expected ", resExpected, " got ", res)
		}
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
