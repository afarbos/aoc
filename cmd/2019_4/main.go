package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/afarbos/aoc/pkg/logging"
	"github.com/afarbos/aoc/pkg/read"
)

const (
	separator             = "-"
	resPasswordCount      = 1099
	resPasswordCountLarge = 710
)

var flagInput string

func init() {
	flag.StringVar(&flagInput, "input", "cmd/2019_4/input", "The input file")
	logging.Flags()
}

func isPasswordValid(password int, largerGroup bool) bool {
	pwd := []byte(strconv.Itoa(password))
	if len(pwd) != 6 {
		return false
	}

	sameAdjacentDigits := false
	for i, digit := range pwd[1:] {
		if digit < pwd[i] {
			return false
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
	return sameAdjacentDigits
}

func passwordCount(start, end int, largerGroup bool) int {
	res := 0
	for i := start; i < end; i++ {
		if isPasswordValid(i, largerGroup) {
			res++
		}
	}
	return res
}

func main() {
	inputRange := read.Read(flagInput, separator)

	pwdCount := passwordCount(inputRange[0], inputRange[1], false)
	if pwdCount != resPasswordCount {
		log.Fatal("Expected ", resPasswordCount, " got ", pwdCount)
	} else {
		log.Println(pwdCount)
	}
	pwdCount = passwordCount(inputRange[0], inputRange[1], true)
	if pwdCount != resPasswordCountLarge {
		log.Fatal("Expected ", resPasswordCountLarge, " got ", pwdCount)
	} else {
		log.Println(passwordCount(inputRange[0], inputRange[1], true))
	}
}
