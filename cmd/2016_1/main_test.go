package main

import (
	"strings"
	"testing"

	"github.com/afarbos/aoc/pkg/point"
	"github.com/afarbos/aoc/pkg/test"
	"github.com/afarbos/aoc/pkg/utils"
)

func TestBlockDistance(t *testing.T) {
	const (
		sequences1 = "R2, L3"
		sequences2 = "R2, R2, R2"
		sequences3 = "R5, L5, R5, R3"
	)

	testData := map[string]int{
		sequences1: 5,
		sequences2: 2,
		sequences3: 12,
	}

	for sequences, resExpected := range testData {
		test.EqualInt(t, point.Distance(newPerson(strings.Split(sequences, separator)).position), resExpected)
	}
}

func TestFirsCross(t *testing.T) {
	const (
		sequences   = "R8, R4, R4, R8"
		resExpected = 4
	)

	test.EqualInt(t, point.Distance(newPerson(strings.Split(sequences, separator)).firstCross), resExpected)
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
