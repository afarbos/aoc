package main

import (
	"flag"
	"log"
	"strings"

	"github.com/afarbos/aoc/pkg/convert"
	"github.com/afarbos/aoc/pkg/mathematic"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	dimSeparator              = "x"
	resTotalRibbon            = 3812909
	resTotalWrappingPaperArea = 1598415
	separator                 = "\n"
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func wrappingPaperArea(giftSize string) int {
	dims := strings.Split(giftSize, dimSeparator)
	if len(dims) < 3 {
		return 0
	}

	sizes, err := convert.Atoi(dims...)
	if err != nil {
		log.Fatal(err)
	}

	w, h, l := sizes[0], sizes[1], sizes[2]

	side1, side2, side3 := w*h, h*l, l*w

	return 2*side1 + 2*side2 + 2*side3 + mathematic.MinInt(side1, side2, side3)
}

func totalWrappingPaperArea(giftSizes []string) int {
	return mathematic.SumString(wrappingPaperArea, giftSizes...)
}

func ribbon(giftSize string) int {
	dims := strings.Split(giftSize, dimSeparator)
	if len(dims) < 3 {
		return 0
	}

	sizes, err := convert.Atoi(dims...)
	if err != nil {
		log.Fatal(err)
	}

	w, h, l := sizes[0], sizes[1], sizes[2]

	return 2*mathematic.MinInt(l+w, w+h, h+l) + l*w*h
}

func totalRibbon(giftSizes []string) int {
	return mathematic.SumString(ribbon, giftSizes...)
}

func main() {
	flag.Parse()

	giftSizes := read.Strings(flagInput, separator)
	utils.AssertEqual(totalWrappingPaperArea(giftSizes), resTotalWrappingPaperArea)
	utils.AssertEqual(totalRibbon(giftSizes), resTotalRibbon)
}
