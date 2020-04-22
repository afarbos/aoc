package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/utils"
)

const (
	decrease    = 223450
	fourSuffix  = 123444
	nodouble    = 123789
	oneTwo      = 111122
	oneTwoThree = 112233
	ones        = 111111
)

var (
	testDataPwdValid = map[int]bool{
		decrease:    false,
		fourSuffix:  true,
		nodouble:    false,
		oneTwo:      true,
		oneTwoThree: true,
		ones:        true,
	}
	testDataPwdValidLarge = map[int]bool{
		decrease:    false,
		fourSuffix:  false,
		nodouble:    false,
		oneTwo:      true,
		oneTwoThree: true,
		ones:        false,
	}
)

func TestIsPasswordValid(t *testing.T) {
	for pwd, expectedRes := range testDataPwdValid {
		if isValid, _ := isPasswordValid(pwd, false); isValid != expectedRes {
			t.Error("Is password valid expected", expectedRes, "got", pwd)
		}
	}
}

func TestIsPasswordValidLarge(t *testing.T) {
	for pwd, expectedRes := range testDataPwdValidLarge {
		if isValid, _ := isPasswordValid(pwd, true); isValid != expectedRes {
			t.Error("Is password large valid expected", expectedRes, "got", pwd)
		}
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)
	for i := 0; i < b.N; i++ {
		main()
	}
}
