package main

import (
	"testing"

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
		if hash := findHashCoin(secretKey, fiveZero); hash != resExpected {
			t.Error("expected ", resExpected, " got ", hash)
		}
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
