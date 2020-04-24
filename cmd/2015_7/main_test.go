package main

import (
	"strings"
	"testing"

	"github.com/afarbos/aoc/pkg/utils"
)

const testDataInstructions = `
123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
`

var (
	testDataWireSignal = map[string]map[string]uint16{
		testDataInstructions: {
			"d": 72,
			"e": 507,
			"f": 492,
			"g": 114,
			"h": 65412,
			"i": 65079,
			"x": 123,
			"y": 456,
		},
	}
)

func TestWireSignal(t *testing.T) {
	for instructions, wires := range testDataWireSignal {
		for wire, resExpected := range wires {
			if res := wireSignal(strings.Split(instructions, separator), wire); res != resExpected {
				t.Error("wire ", wire, " expected ", resExpected, " got ", res)
			}
		}
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
