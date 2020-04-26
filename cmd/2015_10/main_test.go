package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
	"github.com/afarbos/aoc/pkg/utils"
)

var (
	testData = map[string]int{
		"1":      2,
		"11":     2,
		"21":     4,
		"1211":   6,
		"111221": 6,
	}
)

func TestLookAndSayLen(t *testing.T) {
	for start, resExpected := range testData {
		test.EqualInt(t, lookAndSayLen(start, 1), resExpected)
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
