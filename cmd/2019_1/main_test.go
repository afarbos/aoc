package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
	"github.com/afarbos/aoc/pkg/utils"
)

func TestRequiredFuel(t *testing.T) {
	testDataRequiredFuel := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}
	for mass, expectedFuel := range testDataRequiredFuel {
		test.EqualInt(t, requiredFuel(mass), expectedFuel)
	}
}

func TestRequiredFuelTotal(t *testing.T) {
	testDataRequiredFuelTotal := map[int]int{
		14:     2,
		1969:   966,
		100756: 50346,
	}
	for mass, expectedFuel := range testDataRequiredFuelTotal {
		test.EqualInt(t, requiredFuelTotal(mass), expectedFuel)
	}
}

func TestSumFuelRequirements(t *testing.T) {
	test.EqualInt(t, sumFuelRequirements([]int{12, 14, 1969, 100756}, requiredFuel), 34241)
	test.EqualInt(t, sumFuelRequirements([]int{14, 1969, 100756}, requiredFuelTotal), 51314)
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
