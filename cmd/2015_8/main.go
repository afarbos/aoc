package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/afarbos/aoc/pkg/mathematic"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

const (
	separator         = "\n"
	resDiffStrUnquote = 1333
	resDiffStrQuote   = 2046
)

func init() {
	utils.Init(&flagInput)
}

func diffStrUnquote(str string) int {
	s, err := strconv.Unquote(str)
	if err != nil {
		log.Fatal(err)
	}
	return len(str) - len(s)
}

func diffStrQuote(str string) int {
	return len(strconv.Quote(str)) - len(str)
}

func main() {
	flag.Parse()
	strings := read.Strings(flagInput, separator)
	utils.AssertEqual(mathematic.SumString(diffStrUnquote, strings...), resDiffStrUnquote)
	utils.AssertEqual(mathematic.SumString(diffStrQuote, strings...), resDiffStrQuote)
}
