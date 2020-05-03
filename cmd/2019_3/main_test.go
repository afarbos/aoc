package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	example1Wire1 = "R8,U5,L5,D3"
	example1Wire2 = "U7,R6,D4,L4"
	example2Wire1 = "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	example2Wire2 = "U62,R66,U55,R34,D71,R55,D58,R83"
	example3Wire1 = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	example3Wire2 = "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
)

func TestClosestManhattanDistance(t *testing.T) {
	const (
		manhattanDistance1 = 6
		manhattanDistance2 = 159
		manhattanDistance3 = 135
	)

	testDataClosestManhattanDistance := map[int][]string{
		manhattanDistance1: {example1Wire1, example1Wire2},
		manhattanDistance2: {example2Wire1, example2Wire2},
		manhattanDistance3: {example3Wire1, example3Wire2},
	}

	for manhattanDistance, wires := range testDataClosestManhattanDistance {
		g := newGrid(wires)

		test.EqualInt(t, g.closestManhattanDistance(), manhattanDistance)
	}
}

func TestClosestIntersection(t *testing.T) {
	const (
		closestIntersection1 = 30
		closestIntersection2 = 610
		closestIntersection3 = 410
	)

	testDataClosestIntersection := map[int][]string{
		closestIntersection1: {example1Wire1, example1Wire2},
		closestIntersection2: {example2Wire1, example2Wire2},
		closestIntersection3: {example3Wire1, example3Wire2},
	}

	for manhattanDistance, wires := range testDataClosestIntersection {
		g := newGrid(wires)

		test.EqualInt(t, g.closestIntersection(), manhattanDistance)
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
