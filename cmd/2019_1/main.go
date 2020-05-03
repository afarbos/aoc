package main

import (
	"flag"
	"math"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func requiredFuel(mass int) int {
	if mass < 9 {
		return 0
	}

	return mass/3 - 2
}

func requiredFuelTotal(mass int) int {
	var (
		fuel      int
		localMass = mass
	)

	for ok := true; ok; ok = localMass > 0 {
		localMass = int(math.Max(float64(requiredFuel(localMass)), 0.0))
		fuel += localMass
	}

	return fuel
}

func sumFuelRequirements(masses []int, f func(int) int) int {
	var totalFuel int
	for _, mass := range masses {
		totalFuel += f(mass)
	}

	return totalFuel
}

func main() {
	const (
		resSumFuelRequirements      = 3420719
		resTotalSumFuelRequirements = 5128195
	)

	flag.Parse()

	masses := read.Read(flagInput)
	utils.AssertEqual(sumFuelRequirements(masses, requiredFuel), resSumFuelRequirements)
	utils.AssertEqual(sumFuelRequirements(masses, requiredFuelTotal), resTotalSumFuelRequirements)
}
