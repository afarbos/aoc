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

	testDataCountNiceString = map[string]int{
		"ugknbfddgicrmopn": 1,
		"aaa":              1,
		"jchzalrnumimnmhp": 0,
		"haegwjzuvuyypxyu": 0,
		"dvszwmarrgswjxmb": 0,
	}

	testDataCountReallyNiceString = map[string]int{
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

func TestCountNiceString(t *testing.T) {
	for str, resExpected := range testDataCountNiceString {
		if count := countNiceString([]string{str}); count != resExpected {
			t.Error(str, " expected ", resExpected, " got ", count)
		}
	}
}

func TestCountReallyNiceString(t *testing.T) {
	for str, resExpected := range testDataCountReallyNiceString {
		if count := countReallyNiceString([]string{str}); count != resExpected {
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
