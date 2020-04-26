package main

import (
	"testing"

	"github.com/afarbos/aoc/pkg/utils"
)

func TestSumAllNumbers(t *testing.T) {
	var testData = map[int]interface{}{
		-1: map[string]interface{}{"a": []interface{}{-2.0, 1.0}},
		-2: []interface{}{-3.0, map[string]interface{}{"a": 1.0}},
		0:  map[string]interface{}{},
		2:  map[string]interface{}{"a": map[string]interface{}{"b": 4.0}, "c": -2.0},
		3:  []interface{}{[]interface{}{[]interface{}{3.0}}},
		4:  []interface{}{1.0, map[string]interface{}{"c": red, "b": 2.0}, 3.0},
		5:  []interface{}{1.0, red, 4.0},
		6:  []interface{}{1.0, 2.0, 3.0},
		7:  map[string]interface{}{"a": 2.0, "b": 5.0},
	}

	for resExpected, i := range testData {
		if res := sumAllNumbers(i, true); res != resExpected {
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
