package main

import (
	"flag"
	"fmt"

	"github.com/afarbos/aoc/pkg/mathematic"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

const (
	separator      = "\n"
	distanceFormat = "%s to %s = %d"

	resShortest = 117
	resLongest  = 909
)

func init() {
	utils.Init(&flagInput)
}

func Perm(set map[string]struct{}, f func([]string)) {
	var (
		list = make([]string, len(set))
		i    = 0
	)

	for v := range set {
		list[i] = v
		i++
	}

	perm(list, f, 0)
}

func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}

	perm(a, f, i+1)

	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func fDistances(directions []string, f func(...int) int) int {
	cityGraph, cities := make(map[string]map[string]int), make(map[string]struct{})

	for _, direction := range directions {
		var (
			src, dst string
			distance int
		)

		if n, err := fmt.Sscanf(direction, distanceFormat, &src, &dst, &distance); err == nil && n == 3 {
			if _, ok := cityGraph[dst]; !ok {
				cityGraph[dst] = make(map[string]int)
			}

			if _, ok := cityGraph[src]; !ok {
				cityGraph[src] = make(map[string]int)
			}

			cities[dst], cities[src] = struct{}{}, struct{}{}
			cityGraph[dst][src], cityGraph[src][dst] = distance, distance
		}
	}

	var res int = f()

	Perm(cities, func(cities []string) {
		var sum int
		for i, city := range cities[1:] {
			sum += cityGraph[cities[i]][city]
		}
		res = f(res, sum)
	})

	return res
}

func main() {
	flag.Parse()

	directions := read.Strings(flagInput, separator)
	utils.AssertEqual(fDistances(directions, mathematic.MinInt), resShortest)
	utils.AssertEqual(fDistances(directions, mathematic.MaxInt), resLongest)
}
