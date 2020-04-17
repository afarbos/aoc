package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

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

func minInt(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}

func totalWrappingPaperArea(giftSizes []string) int {
	sum := 0
	for _, giftSize := range giftSizes {
		dims := strings.Split(giftSize, dimSeparator)
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
		sum += 2*side1 + 2*side2 + 2*side3 + minInt(side1, side2, side3)
	}
	return sum
}

func totalRibbon(giftSizes []string) int {
	sum := 0
	for _, giftSize := range giftSizes {
		dims := strings.Split(giftSize, dimSeparator)
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
		sum += 2*minInt(l+w, w+h, h+l) + l*w*h
	}
	return sum
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
