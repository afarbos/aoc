package utils

import (
	"testing"
)

func TestUtils(_ *testing.T) {
	var s string

	Init(&s)
	EnableTestMain(&s)
	AssertEqual(0, 0)
}
