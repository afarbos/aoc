package intcode

import (
	"log"

	"github.com/afarbos/aoc/pkg/convert"
)

// OpCodes enumeration of operation code.
const (
	// Add paramater 1 and 2 and store it in 3
	Add = iota + 1
	// Multiply paramater 1 and 2 and store it in 3
	Multiply
	// Input stored at parameter 1
	Input
	// Output the value of parameter 1
	Output
	// JumpTrue at parameter 2 if parameter 1 is non-zero
	JumpTrue
	// JumpFalse at parameter 2 if parameter 2 is zero
	JumpFalse
	// Less return 1 if parameter 1 is less than 2 else 0
	Less
	// Equals return 1 if parameter 1 is equal to 2 else 0
	Equals
	// Stop the program
	Stop = 99
)

// Option to run the intcode computer including the max opcode or the input for a program.
type Option struct {
	index, Input, MaxOp, next int
	res                       []int
}

func extractArguments(instructions []int, index int) (int, int, int, int) {
	var (
		instruction = instructions[index]
		arg1        = index + 1
		arg2        = index + 2
		arg3        = index + 3
		a           = instruction % 100000 / 10000
		b           = instruction % 10000 / 1000
		c           = instruction % 1000 / 100
		op          = instruction % 100
	)

	if a == 0 && arg3 < len(instructions) {
		arg3 = instructions[arg3]
	}

	if b == 0 && arg2 < len(instructions) {
		arg2 = instructions[arg2]
	}

	if c == 0 && arg1 < len(instructions) {
		arg1 = instructions[arg1]
	}

	return op, arg1, arg2, arg3
}

func computeInstruction(instructions []int, opt *Option) bool {
	op, arg1, arg2, arg3 := extractArguments(instructions, opt.index)

	if op != Stop && op > opt.MaxOp {
		return false
	}

	switch op {
	case Add:
		instructions[arg3] = instructions[arg1] + instructions[arg2]
	case Multiply:
		instructions[arg3] = instructions[arg1] * instructions[arg2]
	case Input:
		instructions[arg1] = opt.Input
		opt.next = 2
	case Output:
		opt.res = append(opt.res, instructions[arg1])
		opt.next = 2
	case JumpTrue:
		opt.next = 3
		if instructions[arg1] != 0 {
			opt.index = instructions[arg2]
			opt.next = 0
		}
	case JumpFalse:
		opt.next = 3
		if instructions[arg1] == 0 {
			opt.index = instructions[arg2]
			opt.next = 0
		}
	case Less:
		instructions[arg3] = convert.Btoi(instructions[arg1] < instructions[arg2])
	case Equals:
		instructions[arg3] = convert.Btoi(instructions[arg1] == instructions[arg2])
	case Stop:
		return true
	}

	return false
}

// Compute run a set instructions (also called program) using the Option provided.
func Compute(instructions []int, opt *Option) int {
	for opt.index = 0; opt.index < len(instructions); opt.index += opt.next {
		opt.next = 4
		if computeInstruction(instructions, opt) {
			if len(opt.res) > 0 {
				return opt.res[len(opt.res)-1]
			}

			return instructions[0]
		}
	}
	log.Fatal("No stop code found")

	return 0
}
