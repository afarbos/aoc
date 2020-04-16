package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/logging"
)

var computeInstructions = map[int][]int{
	2:    {1, 0, 0, 0, 99},
	3:    {3, 3, 0, 3, 99},
	4:    {4, 4, 4, 5, 99, 0},
	30:   {1, 1, 1, 4, 99, 5, 6, 0, 99},
	3500: {1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
}

func TestCompute(t *testing.T) {
	for expectedRes, inst := range computeInstructions {
		res := compute(inst, inst[1], inst[2])
		if res != expectedRes {
			t.Error("compute expected", expectedRes, "got", res)
		}
	}
}

func BenchmarkMain(b *testing.B) {
	flagInput = "input"
	logging.Disable()
	for i := 0; i < b.N; i++ {
		main()
	}
}
