package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	comet  = "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds."
	dancer = "Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds."
)

func TestMaxDistance(t *testing.T) {
	var (
		testData = map[string]int{
			comet:  1120,
			dancer: 1056,
		}
		duration = 1000
	)

	for reindeer, resExpected := range testData {
		test.EqualInt(t, maxDistance([]string{reindeer}, duration), resExpected)
	}
}

func TestMaxPoint(t *testing.T) {
	var (
		testData = []string{comet, dancer}
		duration = 1000
		res      = 689
	)

	test.EqualInt(t, maxPoint(testData, duration), res)
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
