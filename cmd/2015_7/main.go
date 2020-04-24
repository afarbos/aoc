package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

const (
	separator = "\n"
	base10    = 10
	bitSize16 = 16

	overrideSignal       = "b"
	resWireA             = 3176
	resWireAWithOverride = 14710
	retWire              = "a"

	mov     = "%s -> %s"
	not     = "%s %s -> %s"
	generic = "%s %s %s -> %s"

	opMov    = ""
	opNot    = "NOT"
	opAnd    = "AND"
	opOr     = "OR"
	opLshift = "LSHIFT"
	opRshift = "RSHIFT"
)

type Wires struct {
	wires map[string]Wire
	cache map[string]uint16
}

type Wire struct {
	op, src1, src2 string
}

func init() {
	utils.Init(&flagInput)
}

func newWires() *Wires {
	return &Wires{make(map[string]Wire), make(map[string]uint16)}
}

func (w *Wires) Add(dst, op, src1, src2 string) {
	w.wires[dst] = Wire{op, src1, src2}
}

func (w *Wires) signalValue(key string) uint16 {
	var res uint16

	if res, err := strconv.ParseUint(key, base10, bitSize16); err == nil {
		return uint16(res)
	} else if res, ok := w.cache[key]; ok {
		return res
	}

	currentWire, ok := w.wires[key]
	if !ok {
		log.Fatal("Unknown wire", key)
	}

	switch currentWire.op {
	case opMov:
		res = w.signalValue(currentWire.src1)
	case opAnd:
		res = w.signalValue(currentWire.src1) & w.signalValue(currentWire.src2)
	case opOr:
		res = w.signalValue(currentWire.src1) | w.signalValue(currentWire.src2)
	case opLshift:
		res = w.signalValue(currentWire.src1) << w.signalValue(currentWire.src2)
	case opRshift:
		res = w.signalValue(currentWire.src1) >> w.signalValue(currentWire.src2)
	case opNot:
		res = ^w.signalValue(currentWire.src1)
	default:
		log.Fatal("Unknown operator", currentWire.op)
	}

	w.cache[key] = res

	return res
}

func wireSignal(instructions []string, key ...string) uint16 {
	wires, ret := newWires(), retWire
	if len(key) > 0 {
		ret = key[0]
	}

	for _, instruction := range instructions {
		var op, src1, src2, dst string

		if instruction == "" {
			continue
		} else if n, err := fmt.Sscanf(instruction, mov, &src1, &dst); err == nil && n == 2 {
		} else if n, err := fmt.Sscanf(instruction, not, &op, &src1, &dst); err == nil && n == 3 {
		} else if n, err := fmt.Sscanf(instruction, generic, &src1, &op, &src2, &dst); err == nil && n == 4 {
		} else {
			log.Fatal("Unknown instruction: ", instruction)
		}

		wires.Add(dst, op, src1, src2)
	}

	return wires.signalValue(ret)
}

func main() {
	flag.Parse()

	var (
		instructions = read.Strings(flagInput, separator)
		signal       = wireSignal(instructions)
	)

	utils.AssertEqual(int(signal), resWireA)

	newInstructions := append(instructions, fmt.Sprintf(mov, strconv.FormatUint(uint64(signal), base10), overrideSignal))
	utils.AssertEqual(int(wireSignal(newInstructions)), resWireAWithOverride)
}
