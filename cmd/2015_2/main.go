package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

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
	w, err := strconv.Atoi(dims[0])
	if err != nil {
		log.Fatal(err)
	}

	h, err := strconv.Atoi(dims[1])
	if err != nil {
		log.Fatal(err)
	}

	l, err := strconv.Atoi(dims[2])
	if err != nil {
		log.Fatal(err)
	}

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
	w, err := strconv.Atoi(dims[0])
	if err != nil {
		log.Fatal(err)
	}

	h, err := strconv.Atoi(dims[1])
	if err != nil {
		log.Fatal(err)
	}

	l, err := strconv.Atoi(dims[2])
	if err != nil {
		log.Fatal(err)
	}
	return 2*mathematic.MinInt(l+w, w+h, h+l) + l*w*h

}

func totalRibbon(giftSizes []string) int {
	return mathematic.SumString(ribbon, giftSizes...)
}

func main() {
	flag.Parse()
	giftSizes := read.Strings(flagInput, separator)
	if wrappingPaperArea := totalWrappingPaperArea(giftSizes); wrappingPaperArea != resTotalWrappingPaperArea {
		log.Fatal("total wrapping paper area expected ", resTotalWrappingPaperArea, " got ", wrappingPaperArea)
	} else {
		log.Println(wrappingPaperArea)
	}
	if ribbon := totalRibbon(giftSizes); ribbon != resTotalRibbon {
		log.Fatal("total ribbon expected ", resTotalRibbon, " got ", ribbon)
	} else {
		log.Println(ribbon)
	}
}
