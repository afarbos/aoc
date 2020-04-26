package test

import "testing"

// EqualBool use != operator to compare the arguments, then log and fail if ture.
func EqualBool(t *testing.T, res, resExpected bool) {
	if res != resExpected {
		t.Error("Expected ", resExpected, " got ", res)
	}
}

// EqualInt use != operator to compare the arguments, then log and fail if ture.
func EqualInt(t *testing.T, res, resExpected int) {
	if res != resExpected {
		t.Error("Expected ", resExpected, " got ", res)
	}
}

// EqualString use != operator to compare the arguments, then log and fail if ture.
func EqualString(t *testing.T, res, resExpected string) {
	if res != resExpected {
		t.Error("Expected ", resExpected, " got ", res)
	}
}

// EqualUint16 use != operator to compare the arguments, then log and fail if ture.
func EqualUint16(t *testing.T, res, resExpected uint16) {
	if res != resExpected {
		t.Error("Expected ", resExpected, " got ", res)
	}
}
