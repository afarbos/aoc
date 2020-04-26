package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	emptyString = `""`
	abc         = `"abc"`
	sixA        = `"aaa\"aaa"`
	hexa        = `"\x27"`
)

var (
	testDataDiffStrUnquote = map[string]int{
		emptyString: 2,
		abc:         2,
		sixA:        3,
		hexa:        5,
	}
	testDataDiffStrQuote = map[string]int{
		emptyString: 4,
		abc:         4,
		sixA:        6,
		hexa:        5,
	}
)

func TestDiffStrUnquote(t *testing.T) {
	for str, resExpected := range testDataDiffStrUnquote {
		test.EqualInt(t, diffStrUnquote(str), resExpected)
	}
}

func TestDiffStrQuote(t *testing.T) {
	for str, resExpected := range testDataDiffStrQuote {
		test.EqualInt(t, diffStrQuote(str), resExpected)
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
