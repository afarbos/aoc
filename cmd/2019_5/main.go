package main

import (
	"flag"

	"github.com/afarbos/aoc/pkg/intcode"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func main() {
	const (
		inputAC      = 1
		inputTR      = 5
		resComputeAC = 9938601
		resComputeTR = 4283952
		separator    = ","
	)

	flag.Parse()

	var (
		instructionsAC = read.Read(flagInput, separator)
		instructionsTR = make([]int, len(instructionsAC))
		optionAC       = &intcode.Option{MaxOp: intcode.Output, Input: inputAC}
		optionTR       = &intcode.Option{MaxOp: intcode.Equals, Input: inputTR}
	)

	copy(instructionsTR, instructionsAC)

	utils.AssertEqual(intcode.Compute(instructionsAC, optionAC), resComputeAC)
	utils.AssertEqual(intcode.Compute(instructionsTR, optionTR), resComputeTR)
}
