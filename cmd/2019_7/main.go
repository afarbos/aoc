package main

import (
	"flag"

	"github.com/afarbos/aoc/pkg/intcode"
	"github.com/afarbos/aoc/pkg/mathematic"
	"github.com/afarbos/aoc/pkg/perm"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func compute(instructions []int, inputs []chan int, amplifier, nextAmplifier int, done chan struct{}) {
	localInstructions := make([]int, len(instructions))
	copy(localInstructions, instructions)
	intcode.Compute(localInstructions, intcode.Equals, inputs[amplifier], inputs[nextAmplifier])
	done <- struct{}{}
}

func findMaxThrusterSignalWithSettings(instructions []int, settingsSet map[int]struct{}) int {
	var signals []int

	perm.Ints(settingsSet, func(settings []int) {
		var (
			inputs = make([]chan int, 0, len(settings))
			done   = make(chan struct{})
		)

		for range settings {
			inputs = append(inputs, make(chan int, 1))
		}

		for amplifier, phase := range settings {
			go compute(instructions, inputs, amplifier, (amplifier+1)%len(settings), done)
			inputs[amplifier] <- phase
		}
		inputs[0] <- 0

		for range settings {
			<-done
		}

		signals = append(signals, <-inputs[0])
	})

	return mathematic.MaxInt(signals...)
}

func findMaxThrusterSignal04(instructions []int) int {
	settingsSet := map[int]struct{}{
		0: {},
		1: {},
		2: {},
		3: {},
		4: {},
	}

	return findMaxThrusterSignalWithSettings(instructions, settingsSet)
}

func findMaxThrusterSignal59(instructions []int) int {
	settingsSet := map[int]struct{}{
		5: {},
		6: {},
		7: {},
		8: {},
		9: {},
	}

	return findMaxThrusterSignalWithSettings(instructions, settingsSet)
}

func main() {
	const (
		resMaxThrusterSignal04 = 116680
		resMaxThrusterSignal59 = 89603079
		separator              = ","
	)

	flag.Parse()

	instructions := read.Read(flagInput, separator)

	utils.AssertEqual(findMaxThrusterSignal04(instructions), resMaxThrusterSignal04)
	utils.AssertEqual(findMaxThrusterSignal59(instructions), resMaxThrusterSignal59)
}
