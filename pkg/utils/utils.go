package utils

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/afarbos/aoc/pkg/logging"
)

const (
	cmd   = "cmd"
	input = "input"
)

// Init add flag input and enable log flags.
func Init(inputFlag *string) {
	flag.StringVar(inputFlag, input, path.Join(cmd, os.Args[0], input), "The input file")
	logging.Flags()
}

// EnableTestMain set inputFlag to input and disable log.
func EnableTestMain(inputFlag *string) {
	*inputFlag = input

	logging.Disable()
}

// AssertEqual verify that the argument are equal, log them and panic if they are different.
func AssertEqual(res, expected int) {
	if res != expected {
		log.Fatal("Expected ", expected, " got ", res)
	} else {
		log.Println(res)
	}
}
