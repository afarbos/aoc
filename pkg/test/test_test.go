package test

import (
	"testing"
)

func TestEqual(t *testing.T) {
	EqualBool(t, true, true)
	EqualInt(t, 0, 0)
	EqualString(t, "", "")
	EqualUint16(t, uint16(1), uint16(1))
}
