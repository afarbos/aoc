package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/logging"
)

func TestRequiredFuel(t *testing.T) {
	testData := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}
	for mass, expectedFuel := range testData {
		if fuel := requiredFuel(mass); fuel != expectedFuel {
			t.Error("fuel: expected", expectedFuel, "got", fuel, "mass", mass)
		}
	}
}

func TestRequiredFuelTotal(t *testing.T) {
	testData := map[int]int{
		14:     2,
		1969:   966,
		100756: 50346,
	}
	for mass, expectedFuel := range testData {
		if fuel := requiredFuelTotal(mass); fuel != expectedFuel {
			t.Error("total fuel: expected", expectedFuel, "got", fuel, "mass", mass)
		}
	}
}

func TestSumFuelRequirements(t *testing.T) {
	if sumFuelRequirements([]int{12, 14, 1969, 100756}, requiredFuel) != 34241 {
		t.Error("sum fuel requirements")
	}
	if sumFuelRequirements([]int{14, 1969, 100756}, requiredFuelTotal) != 51314 {
		t.Error("sum fuel requirements total")
	}
}

func BenchmarkMain(b *testing.B) {
	flagInput = "input"
	logging.Disable()
	for i := 0; i < b.N; i++ {
		main()
	}
}
