package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/afarbos/aoc/pkg/mathematic"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

type reindeer struct {
	name                         string
	duration, point, rest, speed int
}

const (
	raceDuration = 2503
	reindeerLine = "%s can fly %d km/s for %d seconds, but then must rest for %d seconds."
	resDistance  = 2696
	resPoint     = 1084
	separator    = "\n"
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func newReindeer(reindeerStr string) *reindeer {
	r := new(reindeer)
	if n, err := fmt.Sscanf(reindeerStr, reindeerLine, &r.name, &r.speed, &r.duration, &r.rest); err != nil || n != 4 {
		log.Fatalf("error: %s, n: %d", err, n)
	}

	return r
}

func (r *reindeer) distance(duration int) int {
	distance := duration / (r.duration + r.rest) * r.speed * r.duration
	distance += mathematic.MinInt(duration%(r.duration+r.rest), r.duration) * r.speed

	return distance
}

func (r *reindeer) isFlying(duration int) bool {
	return duration%(r.duration+r.rest) < r.duration
}

func maxDistance(reindeers []string, raceDuration int) int {
	res := mathematic.MaxInt()

	for _, reindeerStr := range reindeers {
		if reindeerStr == "" {
			continue
		}

		res = mathematic.MaxInt(res, newReindeer(reindeerStr).distance(raceDuration))
	}

	return res
}

func maxPoint(reindeers []string, raceDuration int) int {
	var reindeersPostitions = make(map[*reindeer]*int)
	for _, reindeerStr := range reindeers {
		if reindeerStr == "" {
			continue
		}
		var d = 0
		reindeersPostitions[newReindeer(reindeerStr)] = &d
	}

	for s := 0; s < raceDuration; s++ {
		var leads []*reindeer
		for r, d := range reindeersPostitions {
			if r.isFlying(s) {
				*d += r.speed
			}
			if len(leads) == 0 || *reindeersPostitions[leads[0]] < *d {
				leads = []*reindeer{r}
			} else if *reindeersPostitions[leads[0]] == *d {
				leads = append(leads, r)
			}
		}

		for _, r := range leads {
			r.point++
		}
	}

	res := mathematic.MaxInt()

	for r := range reindeersPostitions {
		res = mathematic.MaxInt(res, r.point)
	}

	return res
}

func main() {
	flag.Parse()

	reindeers := read.Strings(flagInput, separator)
	utils.AssertEqual(maxDistance(reindeers, raceDuration), resDistance)
	utils.AssertEqual(maxPoint(reindeers, raceDuration), resPoint)
}
