package main

import (
	"flag"

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

	resCountNiceString       = 236
	resCountReallyNiceString = 51
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

func countNiceString(strSlice []string) int {
	count := 0

	for _, str := range strSlice {
		if len(str) < 2 {
			continue
		}

		vowelsCount, charTwiceInRow, containBadStr := 0, false, false

		for pos, char := range str {
			if isVowels(char) {
				vowelsCount++
			}

			if pos != 0 {
				if str[pos] == str[pos-1] {
					charTwiceInRow = true
				}

				if isBadStr(str[pos-1 : pos+1]) {
					containBadStr = true
				}
			}
		}

		if vowelsCount > maxNotNiceVowelsCount && charTwiceInRow && !containBadStr {
			count++
		}
	}

	return count
}

func countReallyNiceString(strSlice []string) int {
	count := 0

	for _, str := range strSlice {
		if len(str) < 3 {
			continue
		}

		lettersPairCount := make(map[string]int)
		hasDoublePair, repeatOneLetterAppart, overlap := false, false, false

		for pos := range str[1:] {
			if overlap {
				overlap = false
				continue
			}

			currentPair := str[pos : pos+2]

			if _, ok := lettersPairCount[currentPair]; !ok {
				lettersPairCount[currentPair] = 1
			} else {
				lettersPairCount[currentPair]++
				if lettersPairCount[currentPair] > 1 {
					hasDoublePair = true
				}
			}

			if pos < len(str)-2 && currentPair == str[pos+1:pos+3] {
				overlap = true
			}

			if pos < len(str)-2 && str[pos] == str[pos+2] {
				repeatOneLetterAppart = true
			}
		}

		if repeatOneLetterAppart && hasDoublePair {
			count++
		}
	}

	return count
}

func main() {
	flag.Parse()

	stringSlice := read.Strings(flagInput, separator)
	utils.AssertEqual(countNiceString(stringSlice), resCountNiceString)
	utils.AssertEqual(countReallyNiceString(stringSlice), resCountReallyNiceString)
}
