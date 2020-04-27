package utils

import (
	"testing"
)

func TestUtils(t *testing.T) {
	var s string

	Init(&s)
	EnableTestMain(&s)
	AssertEqual(0, 0)
}
