package utils

import (
	"flag"
	"os"
	"path"

	"github.com/afarbos/aoc/pkg/logging"
)

const (
	cmd   = "cmd"
	input = "input"
)

func Init(inputFlag *string) {
	flag.StringVar(inputFlag, input, path.Join(cmd, os.Args[0], input), "The input file")
	logging.Flags()
}

func EnableTestMain(inputFlag *string) {
	*inputFlag = input
	logging.Disable()
}
