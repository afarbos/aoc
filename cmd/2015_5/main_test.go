package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/utils"
)

var (
	testDataIsVowels = map[rune]bool{
		a:   true,
		e:   true,
		i:   true,
		o:   true,
		u:   true,
		'w': false,
	}

	testDataCountNice = map[string]int{
		"ugknbfddgicrmopn": 1,
		"aaa":              1,
		"jchzalrnumimnmhp": 0,
		"haegwjzuvuyypxyu": 0,
		"dvszwmarrgswjxmb": 0,
	}

	testDataCountReallyNice = map[string]int{
		"rxexcbwhiywwwwnu": 1,
		"qjhvhtzxzqqjkmpb": 1,
		"xxyxx":            1,
		"uurcxstgmygtbstg": 0,
		"ieodomkazucvgmuy": 0,
	}
)

func TestIsVowels(t *testing.T) {
	for char, resExpected := range testDataIsVowels {
		if res := isVowels(char); res != resExpected {
			t.Error("expected ", resExpected, " got ", res)
		}
	}
}

func TestCountNice(t *testing.T) {
	for str, resExpected := range testDataCountNice {
		if count := countNice([]string{str}); count != resExpected {
			t.Error(str, " expected ", resExpected, " got ", count)
		}
	}
}

func TestCountReallyNice(t *testing.T) {
	for str, resExpected := range testDataCountReallyNice {
		if count := countReallyNice([]string{str}); count != resExpected {
			t.Error(str, " expected ", resExpected, " got ", count)
		}
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
