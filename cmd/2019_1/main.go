package main

import (
	"flag"
	"log"
	"math"

	"github.com/afarbos/aoc/pkg/logging"
	"github.com/afarbos/aoc/pkg/read"
)

const separator = "\n"

var flagInput string

func init() {
	flag.StringVar(&flagInput, "input", "cmd/2019_1/input", "The input file")
	logging.Flags()
}

func requiredFuel(mass int) int {
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
	masses := read.Read(flagInput, separator)
	log.Println(sumFuelRequirements(masses, requiredFuel))
	log.Println(sumFuelRequirements(masses, requiredFuelTotal))
}
