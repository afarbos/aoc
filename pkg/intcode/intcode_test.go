package intcode

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
)

func TestComputeNoInput(t *testing.T) {
	var computeInstructions = map[int][]int{
		1002: {1002, 4, 3, 4, 33},
		1101: {1101, 100, -1, 4, 0},
		2:    {1, 0, 0, 0, 99},
		30:   {1, 1, 1, 4, 99, 5, 6, 0, 99},
		3500: {1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		3:    {3, 3, 0, 3, 99},
		4:    {4, 4, 4, 5, 99, 0},
	}

	for expectedRes, inst := range computeInstructions {
		test.EqualInt(t, Compute(inst, &Option{MaxOp: Multiply}), expectedRes)
	}
}

func TestComputeInput(t *testing.T) {
	input := 42
	test.EqualInt(t, Compute([]int{3, 0, 4, 0, 99}, &Option{MaxOp: Output, Input: input}), input)
}

func TestComputeEquals(t *testing.T) {
	for _, input := range []int{5, 8, 9} {
		resExpected := 0
		if input == 8 {
			resExpected = 1
		}

		for _, inst := range [][]int{{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, {3, 3, 1108, -1, 8, 3, 4, 3, 99}} {
			test.EqualInt(t, Compute(inst, &Option{MaxOp: Equals, Input: input}), resExpected)
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
			test.EqualInt(t, Compute(inst, &Option{MaxOp: Less, Input: input}), resExpected)
		}
	}
}

func TestComputeJump(t *testing.T) {
	var computeInstructions = [][]int{
		{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
	}

	for resExpected, input := range []int{0, 42} {
		for _, inst := range computeInstructions {
			instructions := make([]int, len(inst))
			copy(instructions, inst)
			test.EqualInt(t, Compute(instructions, &Option{MaxOp: JumpFalse, Input: input}), resExpected)
		}
	}
}
