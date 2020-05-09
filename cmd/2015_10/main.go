package main

import (
	"flag"
	"log"
	"strings"

	"github.com/afarbos/aoc/pkg/convert"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	emptyString    = ""
	iteration1     = 40
	iteration2     = 50
	reslookAndSay1 = 492982
	reslookAndSay2 = 6989950
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func lookAndSayNext(s []int) []int {
	var (
		res    = make([]int, 0, len(s)*2)
		c, num int
	)

	for i := 0; i < len(s); i += num {
		c, num = s[i], 1

		for i+num < len(s) && s[i+num] == c {
			num++
		}

		res = append(res, num, c)
	}

	return res
}

func lookAndSayLen(start string, iteration int) int {
	s, err := convert.Atoi(strings.Split(start, emptyString)...)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < iteration; i++ {
		s = lookAndSayNext(s)
	}

	return len(s)
}

func main() {
	flag.Parse()

	lookAndSayStart := read.String(flagInput)
	if lookAndSayStart != "" {
		lookAndSayStart = lookAndSayStart[:len(lookAndSayStart)-1]
	}

	utils.AssertEqual(lookAndSayLen(lookAndSayStart, iteration1), reslookAndSay1)
	utils.AssertEqual(lookAndSayLen(lookAndSayStart, iteration2), reslookAndSay2)
}
