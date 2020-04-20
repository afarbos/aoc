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
	if diff := mathematic.SumString(diffStrUnquote, strings...); diff != resDiffStrUnquote {
		log.Fatal("Expected ", resDiffStrUnquote, " got ", diff)
	} else {
		log.Println(diff)
	}

	if diff := mathematic.SumString(diffStrQuote, strings...); diff != resDiffStrQuote {
		log.Fatal("Expected ", resDiffStrQuote, " got ", diff)
	} else {
		log.Println(diff)
	}
}
