package main

import (
	"flag"
	"log"
	"regexp"

	"github.com/afarbos/aoc/pkg/convert"
	"github.com/afarbos/aoc/pkg/mathematic"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

const (
	gridSize    = 1000
	regexpMatch = `(toggle|turn on|turn off) (\d+),(\d+) through (\d+),(\d+)`
	separator   = "\n"
	toggle      = "toggle"
	turnOff     = "turn off"
	turnOn      = "turn on"

	resCounLitLights     = 569999
	resCountDimLitLights = 17836115
)

func init() {
	utils.Init(&flagInput)
}

func countLitLights(instructions []string) int {
	count, grid := 0, [gridSize][gridSize]bool{}
	re := regexp.MustCompile(regexpMatch)

	for _, instruction := range instructions {
		m := re.FindStringSubmatch(instruction)
		if len(m) < 6 {
			continue
		}

		var f func(bool) bool

		switch m[1] {
		case toggle:
			f = func(b bool) bool { return !b }
		case turnOn:
			f = func(bool) bool { return true }
		case turnOff:
			f = func(bool) bool { return false }
		}

		pos, err := convert.Atoi(m[2:]...)
		if err != nil {
			log.Fatal(err)
		}

		for i := pos[0]; i <= pos[2]; i++ {
			for j := pos[1]; j <= pos[3]; j++ {
				grid[i][j] = f(grid[i][j])
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[j]); j++ {
			if grid[i][j] {
				count++
			}
		}
	}

	return count
}

func countDimLitLights(instructions []string) int {
	count, grid := 0, [gridSize][gridSize]int{}
	re := regexp.MustCompile(regexpMatch)

	for _, instruction := range instructions {
		m := re.FindStringSubmatch(instruction)
		if len(m) < 6 {
			continue
		}

		var f func(int) int

		switch m[1] {
		case toggle:
			f = func(i int) int { return i + 2 }
		case turnOn:
			f = func(i int) int { return i + 1 }
		case turnOff:
			f = func(i int) int { return mathematic.MaxInt(i-1, 0) }
		}

		pos, err := convert.Atoi(m[2:]...)
		if err != nil {
			log.Fatal(err)
		}

		for i := pos[0]; i <= pos[2]; i++ {
			for j := pos[1]; j <= pos[3]; j++ {
				grid[i][j] = f(grid[i][j])
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[j]); j++ {
			count += grid[i][j]
		}
	}

	return count
}

func main() {
	flag.Parse()

	instructions := read.Strings(flagInput, separator)
	utils.AssertEqual(countLitLights(instructions), resCounLitLights)
	utils.AssertEqual(countDimLitLights(instructions), resCountDimLitLights)
}
