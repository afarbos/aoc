package intcode

import "github.com/afarbos/aoc/pkg/convert"

type opcode int

// OpCodes enumeration of operation code.
const (
	// Add paramater 1 and 2 and store it in 3
	Add opcode = iota + 1
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

type program struct {
	maxOp           opcode
	index           int
	inputs, outputs chan int
	runners         []runner
	instructions    []int
}

type runner interface {
	run(p *program, arg1, arg2, arg3 int) bool
}

type add struct{}
type multiply struct{}
type input struct{}
type output struct{}
type jumpTrue struct{}
type jumpFalse struct{}
type less struct{}
type equals struct{}
type stop struct{}

func newProgram(maxOp opcode, inputs, outputs chan int, instructions []int) *program {
	return &program{
		inputs:       inputs,
		instructions: instructions,
		maxOp:        maxOp,
		outputs:      outputs,
		runners: []runner{
			new(stop),
			new(add),
			new(multiply),
			new(input),
			new(output),
			new(jumpTrue),
			new(jumpFalse),
			new(less),
			new(equals),
		},
	}
}

func (p *program) extractArguments() (op opcode, arg1, arg2, arg3 int) {
	var (
		instruction = p.instructions[p.index]
		a           = instruction % 100000 / 10000
		b           = instruction % 10000 / 1000
		c           = instruction % 1000 / 100
	)

	op = opcode(instruction % 100)
	if op > p.maxOp || op == 0 {
		if op != Stop {
			arg1 = -1
		}

		op = 0
	}

	arg3 = p.index + 3
	if a == 0 && arg3 < len(p.instructions) {
		arg3 = p.instructions[arg3]
	}

	arg2 = p.index + 2
	if b == 0 && arg2 < len(p.instructions) {
		arg2 = p.instructions[arg2]
	}

	arg1 = p.index + 1
	if c == 0 && arg1 < len(p.instructions) {
		arg1 = p.instructions[arg1]
	}

	return op, arg1, arg2, arg3
}

func (*add) run(p *program, arg1, arg2, arg3 int) bool {
	p.index += 4
	p.instructions[arg3] = p.instructions[arg1] + p.instructions[arg2]

	return true
}

func (*multiply) run(p *program, arg1, arg2, arg3 int) bool {
	p.index += 4
	p.instructions[arg3] = p.instructions[arg1] * p.instructions[arg2]

	return true
}

func (*input) run(p *program, arg1, _, _ int) bool {
	p.index += 2
	p.instructions[arg1] = <-p.inputs

	return true
}

func (*output) run(p *program, arg1, _, _ int) bool {
	p.index += 2
	p.outputs <- p.instructions[arg1]

	return true
}

func (*jumpTrue) run(p *program, arg1, arg2, _ int) bool {
	p.index += 3
	if p.instructions[arg1] != 0 {
		p.index = p.instructions[arg2]
	}

	return true
}

func (*jumpFalse) run(p *program, arg1, arg2, _ int) bool {
	p.index += 3
	if p.instructions[arg1] == 0 {
		p.index = p.instructions[arg2]
	}

	return true
}

func (*less) run(p *program, arg1, arg2, arg3 int) bool {
	p.index += 4
	p.instructions[arg3] = convert.Btoi(p.instructions[arg1] < p.instructions[arg2])

	return true
}

func (*equals) run(p *program, arg1, arg2, arg3 int) bool {
	p.index += 4
	p.instructions[arg3] = convert.Btoi(p.instructions[arg1] == p.instructions[arg2])

	return true
}

func (*stop) run(p *program, arg1, _, _ int) bool {
	if arg1 < 0 { // unknown opcode
		p.index += 4
		return true
	}

	// normal stop
	if p.maxOp < Output {
		p.outputs <- p.instructions[0]
	}

	return false
}

func (p *program) execute() {
	var (
		shouldContinue   = true
		op               opcode
		arg1, arg2, arg3 int
	)

	for ; shouldContinue; shouldContinue = p.runners[op].run(p, arg1, arg2, arg3) {
		op, arg1, arg2, arg3 = p.extractArguments()
	}
}

// Compute run a set instructions (also called program).
func Compute(instructions []int, maxOp opcode, inputs, outputs chan int) {
	newProgram(maxOp, inputs, outputs, instructions).execute()
}
