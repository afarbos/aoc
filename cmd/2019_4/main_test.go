package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
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
		isValid, _ := isPasswordValid(pwd, false)
		test.EqualBool(t, isValid, expectedRes)
	}
}

func TestIsPasswordValidLarge(t *testing.T) {
	for pwd, expectedRes := range testDataPwdValidLarge {
		isValid, _ := isPasswordValid(pwd, true)
		test.EqualBool(t, isValid, expectedRes)
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
