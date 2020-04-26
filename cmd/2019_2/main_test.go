package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/utils"
)

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
