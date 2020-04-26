package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/utils"
)

func TestNextPassword(t *testing.T) {
	var testData = map[string]string{
		"abcdefgh": "abcdffaa",
		"ghijklmn": "ghjaabcc",
	}

	for pwd, resExpected := range testData {
		bPwd := []byte(pwd)
		if nextPassword(bPwd); string(bPwd) != resExpected {
			t.Error("expected ", resExpected, " got ", string(bPwd))
		}
	}
}

func TestIsPasswordValid(t *testing.T) {
	var testData = map[string]bool{
		"hijklmmn": false,
		"abbceffg": false,
		"abbcegjk": false,
	}

	for pwd, resExpected := range testData {
		if res := isPasswordValid([]byte(pwd)); res != resExpected {
			t.Error("expected ", resExpected, " got ", res)
		}
	}
}

func BenchmarkMain(b *testing.B) {
	utils.EnableTestMain(&flagInput)

	for i := 0; i < b.N; i++ {
		main()
	}
}
