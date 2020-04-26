package main

import (
	"flag"
	"log"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	resSum    = 111754
	resSumRed = 65402
	red       = "red"
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func sumAllNumbers(i interface{}, redRule bool) int {
	res := 0
	switch v := i.(type) {
	case float64:
		res += int(v)
	case []interface{}:
		for _, e := range v {
			res += sumAllNumbers(e, redRule)
		}
	case map[string]interface{}:
		for _, e := range v {
			if str, ok := e.(string); redRule && ok && str == red {
				return 0
			}

			res += sumAllNumbers(e, redRule)
		}
	case string:
	default:
		log.Fatalf("Unexpected type %T, value %s", v, v)
	}

	return res
}

func main() {
	flag.Parse()

	var i interface{}

	read.JSON(flagInput, &i)
	utils.AssertEqual(sumAllNumbers(i, false), resSum)
	utils.AssertEqual(sumAllNumbers(i, true), resSumRed)
}
