package mathematic

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
)

func TestMinInt(t *testing.T) {
	test.EqualInt(t, MinInt(-1, 1), -1)
}

func TestMaxInt(t *testing.T) {
	test.EqualInt(t, MaxInt(-1, 1), 1)
}

func TestSumString(t *testing.T) {
	test.EqualInt(t, SumString(func(string) int { return 42 }), 0)
	test.EqualInt(t, SumString(func(s string) int { return len(s) }, "-"), 1)
}
