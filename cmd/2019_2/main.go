package main

import (
	"flag"
	"log"

	"github.com/afarbos/aoc/pkg/logging"
	"github.com/afarbos/aoc/pkg/read"
)

const (
	instruction1 = 12
	instruction2 = 2
	maxNounVerb  = 100
	output       = 19690720
	separator    = ","
)

var flagInput string

func init() {
	flag.StringVar(&flagInput, "input", "cmd/2019_2/input", "The input file")
	logging.Flags()
}

func compute(instructions []int, noun, verb int) int {
	instructions[1] = noun
	instructions[2] = verb
	for index := 0; index < len(instructions); index += 4 {
		switch instructions[index] {
		case 1: // Add
			instructions[instructions[index+3]] = instructions[instructions[index+1]] + instructions[instructions[index+2]]
		case 2: // Multiply
			instructions[instructions[index+3]] = instructions[instructions[index+1]] * instructions[instructions[index+2]]
		case 99: // Stop
			return instructions[0]
		}
	}
	log.Fatal("No 99 code found")
	return 0
}

func findNounVerb(instructions []int, maxNoun, maxVerb, expectedOutput int) int {
	for noun := 0; noun < maxNoun; noun++ {
		for verb := 0; verb < maxVerb; verb++ {
			tmp := make([]int, len(instructions))
			copy(tmp, instructions)
			if compute(tmp, noun, verb) == expectedOutput {
				return 100*noun + verb
			}
		}
	}
	log.Fatal("Did not find noun and verb")
	return 0
}

func main() {
	instructions := read.Read(flagInput, separator)

	instructions2 := make([]int, len(instructions))
	copy(instructions2, instructions)

	log.Println(compute(instructions, instruction1, instruction2))
	log.Println(findNounVerb(instructions2, maxNounVerb, maxNounVerb, output))
}
