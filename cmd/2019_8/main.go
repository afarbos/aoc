package main

import (
	"flag"
	"log"
	"math"
	"strings"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func count12fewest0(image string, wide, tall int) int {
	const r0, r1, r2 = '0', '1', '2'

	var (
		counts = map[rune]int{r0: math.MaxInt32, r1: 0, r2: 0}
		size   = wide * tall
		rows   = len(image) / size
	)

	for long := 0; long < rows; long++ {
		lCounts := make(map[rune]int, len(counts))

		for _, r := range image[long*size : (long+1)*size] {
			lCounts[r]++
		}

		if counts[r0] > lCounts[r0] {
			counts = lCounts
		}
	}

	return counts[r1] * counts[r2]
}

func getWord(image string, wide, tall int) string {
	const black, white, transparent = '0', '1', '2'

	var (
		size      = wide * tall
		row, word string
		rows      = len(image) / wide * tall
		fmter     = map[byte]string{black: " ", white: "#"}
	)

	for i := 0; i < tall; i++ {
		row = ""

		for j := 0; j < wide; j++ {
			for long := 0; long < rows; long++ {
				index := long*size + i*wide + j
				log.Println(index)

				if r := image[index]; r != transparent {
					row += fmter[r]
				}
			}
		}

		word += row + read.Eol
	}

	return word
}

func main() {
	const (
		res12 = 2159
		wide  = 25
		tall  = 6
	)

	flag.Parse()

	picture := strings.TrimSuffix(read.String(flagInput), read.Eol)

	utils.AssertEqual(count12fewest0(picture, wide, tall), res12)
	log.Println(getWord(picture, wide, tall))
}
