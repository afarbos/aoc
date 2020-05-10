package main

import (
	"container/ring"
	"flag"
	"image"
	"log"
	"strconv"
	"strings"

	"github.com/afarbos/aoc/pkg/point"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

const separator = ", "

var flagInput string

func init() {
	utils.Init(&flagInput)
}

type person struct {
	firstCross, position image.Point
	direction            *ring.Ring
	visited              map[image.Point]struct{}
}

func newPerson(sequences []string) *person {
	var (
		p                                        = new(person)
		northRing, eastRing, southRing, westRing = ring.New(1), ring.New(1), ring.New(1), ring.New(1)
	)

	northRing.Value = image.Point{1, 0}
	eastRing.Value = image.Point{0, 1}
	southRing.Value = image.Point{-1, 0}
	westRing.Value = image.Point{0, -1}

	southRing.Link(westRing)
	eastRing.Link(southRing)
	northRing.Link(eastRing)

	p.direction = northRing
	p.visited = make(map[image.Point]struct{}, len(sequences)*2)

	p.walkBlocks(sequences)

	return p
}

func (p *person) walk(newDirection byte, blockCount int) {
	const (
		right = "R"
		left  = "L"
	)

	switch newDirection {
	case right[0]:
		p.direction = p.direction.Next()
	case left[0]:
		p.direction = p.direction.Prev()
	default:
		log.Fatal(newDirection)
	}

	var (
		direction    = p.direction.Value.(image.Point)
		nextPosition = p.position.Add(direction.Mul(blockCount))
	)

	zp := image.Point{}
	if p.firstCross == zp {
		for pt := p.position; pt != nextPosition; pt = pt.Add(direction) {
			if _, ok := p.visited[pt]; ok {
				p.firstCross = pt
			}

			p.visited[pt] = struct{}{}
		}
	}

	p.position = nextPosition
}

func (p *person) walkBlocks(sequences []string) {
	for i := range sequences {
		blockCount, err := strconv.Atoi(strings.TrimSuffix(sequences[i][1:], read.Eol))
		if err != nil {
			log.Fatal(err)
		}

		p.walk(sequences[i][0], blockCount)
	}
}

func main() {
	const (
		resTotalBlocks = 298
		resFirstCross  = 158
	)

	flag.Parse()

	var (
		sequences = read.Strings(flagInput, separator)
		p         = newPerson(sequences)
	)

	utils.AssertEqual(point.Distance(p.position), resTotalBlocks)
	utils.AssertEqual(point.Distance(p.firstCross), resFirstCross)
}
