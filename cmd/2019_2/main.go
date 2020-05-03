package main

import (
	"flag"
	"log"

	"github.com/afarbos/aoc/pkg/intcode"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func findNounVerb(instructions []int, maxNoun, maxVerb, expectedOutput int) int {
	for noun := 0; noun < maxNoun; noun++ {
		for verb := 0; verb < maxVerb; verb++ {
			tmp := make([]int, len(instructions))
			copy(tmp, instructions)
			tmp[1] = noun
			tmp[2] = verb

			if intcode.Compute(tmp, &intcode.Option{MaxOp: intcode.Multiply}) == expectedOutput {
				return 100*noun + verb
			}
		}
	}
	log.Fatal("Did not find noun and verb")

	return 0
}

func main() {
	const (
		instruction1    = 12
		instruction2    = 2
		maxNounVerb     = 100
		output          = 19690720
		resCompute      = 4090701
		resFindNounVerb = 6421
		separator       = ","
	)

	flag.Parse()

	instructions := read.Read(flagInput, separator)
	instructions2 := make([]int, len(instructions))
	copy(instructions2, instructions)

	instructions[1] = instruction1
	instructions[2] = instruction2
	utils.AssertEqual(intcode.Compute(instructions, &intcode.Option{MaxOp: intcode.Multiply}), resCompute)
	utils.AssertEqual(findNounVerb(instructions2, maxNounVerb, maxNounVerb, output), resFindNounVerb)
}
