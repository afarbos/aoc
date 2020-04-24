package main

import (
	"flag"

	"github.com/afarbos/aoc/pkg/mathematic"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

const (
	separator = "\n"

	a = 'a'
	e = 'e'
	i = 'i'
	o = 'o'
	u = 'u'

	ab = "ab"
	cd = "cd"
	pq = "pq"
	xy = "xy"

	maxNotNiceVowelsCount = 2

	resCountNice       = 236
	resCountReallyNice = 51
)

func init() {
	utils.Init(&flagInput)
}

func isVowels(c rune) bool {
	return c == a || c == e || c == i || c == o || c == u
}

func isBadStr(s string) bool {
	return s == ab || s == cd || s == pq || s == xy
}

func nice(str string) int {
	if len(str) < 2 {
		return 0
	}

	vowelsCount, charTwiceInRow, containBadStr := 0, false, false

	for pos, char := range str {
		if isVowels(char) {
			vowelsCount++
		}

		if pos != 0 {
			charTwiceInRow = charTwiceInRow || str[pos] == str[pos-1]
			containBadStr = containBadStr || isBadStr(str[pos-1:pos+1])
		}
	}

	if vowelsCount > maxNotNiceVowelsCount && charTwiceInRow && !containBadStr {
		return 1
	}

	return 0
}

func countNice(strSlice []string) int {
	return mathematic.SumString(nice, strSlice...)
}

func reallyNice(str string) int {
	if len(str) < 3 {
		return 0
	}

	lettersPairCount := make(map[string]struct{})
	hasDoublePair, repeatOneLetterAppart, overlap := false, false, false

	for pos := range str[1:] {
		if overlap {
			overlap = false
			continue
		}

		currentPair := str[pos : pos+2]

		if _, ok := lettersPairCount[currentPair]; !ok {
			lettersPairCount[currentPair] = struct{}{}
		} else if !hasDoublePair {
			hasDoublePair = true
		}

		overlap = overlap || pos < len(str)-2 && currentPair == str[pos+1:pos+3]
		repeatOneLetterAppart = repeatOneLetterAppart || pos < len(str)-2 && str[pos] == str[pos+2]
	}

	if repeatOneLetterAppart && hasDoublePair {
		return 1
	}

	return 0
}

func countReallyNice(strSlice []string) int {
	return mathematic.SumString(reallyNice, strSlice...)
}

func main() {
	flag.Parse()

	stringSlice := read.Strings(flagInput, separator)
	utils.AssertEqual(countNice(stringSlice), resCountNice)
	utils.AssertEqual(countReallyNice(stringSlice), resCountReallyNice)
}
