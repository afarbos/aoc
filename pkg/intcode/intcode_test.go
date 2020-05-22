package intcode

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
)

func TestComputeNoInput(t *testing.T) {
	computeInstructions := map[int][]int{
		1002: {1002, 4, 3, 4, 33},
		1101: {1101, 100, -1, 4, 0},
		2:    {1, 0, 0, 0, 99},
		30:   {1, 1, 1, 4, 99, 5, 6, 0, 99},
		3500: {1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		3:    {3, 3, 0, 3, 99},
		4:    {4, 4, 4, 5, 99, 0},
	}

	for expectedRes, inst := range computeInstructions {
		outputs := make(chan int, 1)
		Compute(inst, Multiply, make(chan int), outputs)
		test.EqualInt(t, <-outputs, expectedRes)
	}
}

func TestComputeInput(t *testing.T) {
	const input = 42

	inputs, outputs := make(chan int, 1), make(chan int, 1)
	inputs <- input
	Compute([]int{3, 0, 4, 0, 99}, Output, inputs, outputs)
	test.EqualInt(t, <-outputs, input)
}

func TestComputeEquals(t *testing.T) {
	for _, input := range []int{5, 8, 9} {
		resExpected := 0

		if input == 8 {
			resExpected = 1
		}

		for _, inst := range [][]int{{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, {3, 3, 1108, -1, 8, 3, 4, 3, 99}} {
			inputs, outputs := make(chan int, 1), make(chan int, 1)
			inputs <- input
			Compute(inst, Equals, inputs, outputs)
			test.EqualInt(t, <-outputs, resExpected)
		}
	}
}

func TestComputeLess(t *testing.T) {
	for _, input := range []int{5, 8, 9} {
		resExpected := 0

		if input < 8 {
			resExpected = 1
		}

		for _, inst := range [][]int{{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, {3, 3, 1107, -1, 8, 3, 4, 3, 99}} {
			inputs, outputs := make(chan int, 1), make(chan int, 1)
			inputs <- input
			Compute(inst, Less, inputs, outputs)
			test.EqualInt(t, <-outputs, resExpected)
		}
	}
}

func TestComputeJump(t *testing.T) {
	computeInstructions := [][]int{
		{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
	}

	for resExpected, input := range []int{0, 42} {
		for _, inst := range computeInstructions {
			var (
				instructions    = make([]int, len(inst))
				inputs, outputs = make(chan int, 1), make(chan int, 1)
			)

			copy(instructions, inst)
			inputs <- input
			Compute(instructions, JumpFalse, inputs, outputs)

			test.EqualInt(t, <-outputs, resExpected)
		}
	}
}
