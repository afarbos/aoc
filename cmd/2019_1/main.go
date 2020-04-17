package main

import (
	"flag"
	"log"
	"math"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	resSumFuelRequirements      = 3420719
	resTotalSumFuelRequirements = 5128195
	separator                   = "\n"
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
	localMass := mass
	var fuel int
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
	flag.Parse()
	masses := read.Read(flagInput, separator)
	if fuel := sumFuelRequirements(masses, requiredFuel); fuel != resSumFuelRequirements {
		log.Fatal("Total fuel expected ", resSumFuelRequirements, " got ", fuel)
	} else {
		log.Println(fuel)
	}
	if fuel := sumFuelRequirements(masses, requiredFuelTotal); fuel != resTotalSumFuelRequirements {
		log.Fatal("Total fuel expected ", resTotalSumFuelRequirements, " got ", fuel)
	} else {
		log.Println(fuel)
	}
}
