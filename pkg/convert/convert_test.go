package convert

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
)

func TestAtoi(t *testing.T) {
	res, _ := Atoi("1", "e", "2")
	test.EqualInt(t, len(res), 3)
	test.EqualInt(t, res[0], 1)
	test.EqualInt(t, res[2], 2)
}

func TestBtoi(t *testing.T) {
	test.EqualInt(t, Btoi(true), 1)
	test.EqualInt(t, Btoi(false), 0)
}
