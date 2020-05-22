package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/afarbos/aoc/pkg/mathematic"
	"github.com/afarbos/aoc/pkg/perm"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	dot                    = "."
	happinessLine          = "%s would %s %d happiness units by sitting next to %s"
	lose                   = "lose"
	myself                 = "myself"
	resOptimalChange       = 709
	resOptimalChangeWithMe = 668
	separator              = "\n"
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func loadHappinessPerPerson(happinessChanges []string, me bool) map[string]map[string]int {
	happinessPerPerson := make(map[string]map[string]int)

	for _, happinessChange := range happinessChanges {
		var (
			src, dst, change string
			amount           int
		)

		if happinessChange == "" {
			continue
		} else if n, err := fmt.Sscanf(happinessChange, happinessLine, &src, &change, &amount, &dst); err != nil || n != 4 {
			log.Fatalf("error: %s, n: %d", err, n)
		}

		dst = strings.TrimSuffix(dst, dot)

		if _, ok := happinessPerPerson[src]; !ok {
			happinessPerPerson[src] = make(map[string]int, len(happinessPerPerson))
		}

		if change == lose {
			amount = -amount
		}

		happinessPerPerson[src][dst] = amount
	}

	if me {
		happinessPerPerson[myself] = make(map[string]int, len(happinessPerPerson))
		for key := range happinessPerPerson {
			happinessPerPerson[myself][key] = 0
		}
	}

	return happinessPerPerson
}

func optimalHappinessChange(happinessChanges []string, me bool) int {
	happinessPerPerson := loadHappinessPerPerson(happinessChanges, me)
	persons := make(map[string]struct{})

	for person := range happinessPerPerson {
		persons[person] = struct{}{}
	}

	var res = mathematic.MaxInt()

	perm.Strings(persons, func(persons []string) {
		var (
			sum  int
			last = len(persons) - 1
		)

		for i, p := range persons {
			switch i {
			case 0:
				sum += happinessPerPerson[p][persons[i+1]]
				sum += happinessPerPerson[p][persons[last]]
			case last:
				sum += happinessPerPerson[p][persons[i-1]]
				sum += happinessPerPerson[p][persons[0]]
			default:
				sum += happinessPerPerson[p][persons[i-1]]
				sum += happinessPerPerson[p][persons[i+1]]
			}
		}

		res = mathematic.MaxInt(res, sum)
	})

	return res
}

func main() {
	flag.Parse()

	happinessChanges := read.Strings(flagInput, separator)
	utils.AssertEqual(optimalHappinessChange(happinessChanges, false), resOptimalChange)
	utils.AssertEqual(optimalHappinessChange(happinessChanges, true), resOptimalChangeWithMe)
}
