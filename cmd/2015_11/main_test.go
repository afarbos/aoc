package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
	"github.com/afarbos/aoc/pkg/utils"
)

func TestNextPassword(t *testing.T) {
	var testData = map[string]string{
		"abcdefgh": "abcdffaa",
		"ghijklmn": "ghjaabcc",
	}

	for pwd, resExpected := range testData {
		bPwd := []byte(pwd)
		nextPassword(bPwd)
		test.EqualString(t, string(bPwd), resExpected)
	}
}

func TestIsPasswordValid(t *testing.T) {
	var testData = map[string]bool{
		"hijklmmn": false,
		"abbceffg": false,
		"abbcegjk": false,
	}

	for pwd, resExpected := range testData {
		test.EqualBool(t, isPasswordValid([]byte(pwd)), resExpected)
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
