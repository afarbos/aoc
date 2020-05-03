package main

import (
	"flag"
	"image"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/afarbos/aoc/pkg/mathematic"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

type point struct {
	image.Point
}

type wire struct {
	index int
	step  int
}

type grid struct {
	grid  map[point][]wire
	cross []point
}

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func newGrid(wirePaths []string) *grid {
	g := new(grid)
	g.grid = make(map[point][]wire)

	for index, wirePath := range wirePaths {
		if wirePath == "" {
			continue
		}

		g.cross = []point{}
		g.addWire(wirePath, index)
	}

	return g
}

func (g *grid) addWire(wirePath string, index int) {
	const wireSepartor = ","
	directions := strings.Split(wirePath, wireSepartor)

	currentPt := point{image.Pt(0, 0)}
	currentStep := 0
	g.grid[currentPt] = []wire{{index, currentStep}}

	for _, direction := range directions {
		step, err := strconv.Atoi(direction[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch direction[0] {
		case 'R': // Right
			currentStep = g.add(&currentPt, &currentPt.X, step, currentStep, index)
		case 'L': // Left
			currentStep = g.add(&currentPt, &currentPt.X, -step, currentStep, index)
		case 'U': // Up
			currentStep = g.add(&currentPt, &currentPt.Y, step, currentStep, index)
		case 'D': // Down
			currentStep = g.add(&currentPt, &currentPt.Y, -step, currentStep, index)
		}
	}
}

func (g *grid) add(position *point, coordinate *int, step, currentStep, index int) int {
	next := 1
	if step < 0 {
		next = -1
	}

	for i := next; i != step+next; i += next {
		currentStep++

		*coordinate += next

		g.grid[*position] = append(g.grid[*position], wire{index, currentStep})

		if _, ok := g.grid[*position]; !ok {
			continue
		}

		for _, w := range g.grid[*position] {
			if index != w.index {
				g.cross = append(g.cross, *position)
			}
		}
	}

	return currentStep
}

func (g *grid) closestManhattanDistance() int {
	res := math.MaxInt32
	for _, pt := range g.cross {
		res = mathematic.MinInt(res, pt.manhattanDistance())
	}

	return res
}

func (g *grid) closestIntersection() int {
	res := math.MaxInt32

	for _, pt := range g.cross {
		localSteps := map[int]int{}

		for _, wire := range g.grid[pt] {
			if _, ok := localSteps[wire.index]; !ok {
				localSteps[wire.index] = math.MaxInt32
			}

			localSteps[wire.index] = mathematic.MinInt(localSteps[wire.index], wire.step)
		}

		intersection := 0
		for _, lStep := range localSteps {
			intersection += lStep
		}

		res = mathematic.MinInt(res, intersection)
	}

	return res
}

func (pt *point) manhattanDistance() int {
	return abs(pt.X) + abs(pt.Y)
}

func main() {
	const (
		resClosestIntersection      = 65356
		resClosestManhattanDistance = 1337
	)

	flag.Parse()

	wirePaths := read.Strings(flagInput)
	g := newGrid(wirePaths)

	utils.AssertEqual(g.closestManhattanDistance(), resClosestManhattanDistance)
	utils.AssertEqual(g.closestIntersection(), resClosestIntersection)
}
