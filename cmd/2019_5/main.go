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
		inputAC       = 1
		inputTR       = 5
		outputCountAC = 10
		resComputeAC  = 9938601
		resComputeTR  = 4283952
		separator     = ","
	)

	flag.Parse()

	var (
		instructionsAC       = read.Read(flagInput, separator)
		inChanAC, inChanTR   = make(chan int, 1), make(chan int, 1)
		outChanAC, outChanTR = make(chan int, outputCountAC), make(chan int, 1)
		instructionsTR       = make([]int, len(instructionsAC))
	)

	copy(instructionsTR, instructionsAC)

	inChanAC <- inputAC
	intcode.Compute(instructionsAC, intcode.Output, inChanAC, outChanAC)

	for i := 1; i < outputCountAC; i++ {
		<-outChanAC
	}
	utils.AssertEqual(<-outChanAC, resComputeAC)

	inChanTR <- inputTR
	intcode.Compute(instructionsTR, intcode.Equals, inChanTR, outChanTR)
	utils.AssertEqual(<-outChanTR, resComputeTR)
}
