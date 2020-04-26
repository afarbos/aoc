package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
	"github.com/afarbos/aoc/pkg/utils"
)

var (
	testDataTotalWrappingPaperArea = map[int][]string{
		58:   {"2x3x4"},
		43:   {"1x1x10"},
		1402: {"4x23x21"},
	}
	testDataTotalRibbon = map[int][]string{
		34: {"2x3x4"},
		14: {"1x1x10"},
	}
)

func TestTotalWrappingPaperArea(t *testing.T) {
	for resExpected, giftSizes := range testDataTotalWrappingPaperArea {
		test.EqualInt(t, totalWrappingPaperArea(giftSizes), resExpected)
	}
}

func TestTotalRibbon(t *testing.T) {
	for resExpected, giftSizes := range testDataTotalRibbon {
		test.EqualInt(t, totalRibbon(giftSizes), resExpected)
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
