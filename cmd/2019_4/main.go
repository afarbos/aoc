package main

import (
	"flag"
	"strconv"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func isPasswordValid(password int, largerGroup bool) (bool, int) {
	pwd := []byte(strconv.Itoa(password))
	if len(pwd) != 6 {
		return false, 0
	}

	sameAdjacentDigits := false

	for i, digit := range pwd[1:] {
		if digit < pwd[i] {
			divPower := 1
			for j := 0; j < len(pwd)-i-2; j++ {
				divPower *= 10
			}

			return false, (password/divPower+1)*divPower - password - 1
		} else if digit == pwd[i] {
			if largerGroup {
				if (i < 1 || digit != pwd[i-1]) && (i+2 >= len(pwd) || digit != pwd[i+2]) {
					sameAdjacentDigits = true
				}
			} else {
				sameAdjacentDigits = true
			}
		}
	}

	return sameAdjacentDigits, 0
}

func passwordCount(start, end int, largerGroup bool) int {
	res := 0

	for i := start; i < end; i++ {
		isValid, next := isPasswordValid(i, largerGroup)
		if isValid {
			res++
		}

		i += next
	}

	return res
}

func main() {
	const (
		separator             = "-"
		resPasswordCount      = 1099
		resPasswordCountLarge = 710
	)

	flag.Parse()

	inputRange := read.Read(flagInput, separator)
	utils.AssertEqual(passwordCount(inputRange[0], inputRange[1], false), resPasswordCount)
	utils.AssertEqual(passwordCount(inputRange[0], inputRange[1], true), resPasswordCountLarge)
}
