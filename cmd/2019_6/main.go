package main

import (
	"flag"
	"log"
	"strings"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

type vertex struct {
	name     string
	parent   *vertex
	children []*vertex
}

type graph map[string]*vertex

func (v *vertex) checksum() int {
	var checksum int

	if v.parent == nil {
		return checksum
	}

	for tmpV := v; tmpV.parent != nil; tmpV = tmpV.parent {
		checksum++
	}

	return checksum
}

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func newGraph(edges []string) *graph {
	g := make(graph)

	for _, edge := range edges {
		vertices := strings.Split(edge, ")")
		if len(vertices) < 2 {
			continue
		}

		src, dst := vertices[0], vertices[1]

		if _, ok := g[src]; !ok {
			g[src] = &vertex{name: src}
		}

		if _, ok := g[dst]; !ok {
			g[dst] = &vertex{name: dst}
		}

		g[src].children = append(g[src].children, g[dst])
		g[dst].parent = g[src]
	}

	return &g
}

func (g *graph) orbitSum() int {
	sum := 0

	for _, v := range *g {
		sum += v.checksum()
	}

	return sum
}

func (g *graph) minOrbitalTransfer(src string, dst string) int {
	var (
		dstVertex = (*g)[dst]
		parents   = make(map[string]int)
		srcVertex = (*g)[src]
		transfer  int
	)

	for v := srcVertex.parent; v.parent != nil; v = v.parent {
		parents[v.name] = transfer
		transfer++
	}

	transfer = 0

	for v := dstVertex.parent; v.parent != nil; v = v.parent {
		if t, ok := parents[v.name]; ok {
			return t + transfer
		}

		transfer++
	}

	log.Fatal("Orbital Transfer Not Found")

	return 0
}

func main() {
	const (
		resMinOrbitTransfer = 481
		resOrbitSum         = 315757
		san                 = "SAN"
		separator           = "\n"
		you                 = "YOU"
	)

	flag.Parse()

	var (
		edges = read.Strings(flagInput, separator)
		g     = newGraph(edges)
	)

	utils.AssertEqual(g.orbitSum(), resOrbitSum)
	utils.AssertEqual(g.minOrbitalTransfer(you, san), resMinOrbitTransfer)
}
