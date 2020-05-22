package perm

import (
	"testing"

	"github.com/afarbos/aoc/pkg/test"
)

func TestRead(t *testing.T) {
	count := 0

	Strings(map[string]struct{}{
		"a": {},
		"b": {},
	}, func([]string) { count++ })
	test.EqualInt(t, count, 2)
}
