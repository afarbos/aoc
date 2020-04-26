package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
	"github.com/afarbos/aoc/pkg/utils"
)

var (
	testDataFindHashCoinFive = map[string]int{
		"abcdef":  609043,
		"pqrstuv": 1048970,
	}
)

func TestFindHashCoin(t *testing.T) {
	for secretKey, resExpected := range testDataFindHashCoinFive {
		test.EqualInt(t, findHashCoin(secretKey, fiveZero), resExpected)
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
