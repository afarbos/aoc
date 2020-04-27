package test

import "testing"

// EqualBool use != operator to compare the arguments, then log it and fail.
func EqualBool(t *testing.T, res, resExpected bool) {
	if res != resExpected {
		t.Error("Expected ", resExpected, " got ", res)
	}
}

// EqualInt use != operator to compare the arguments, then log it and fail.
func EqualInt(t *testing.T, res, resExpected int) {
	if res != resExpected {
		t.Error("Expected ", resExpected, " got ", res)
	}
}

// EqualString use != operator to compare the arguments, then log it and fail.
func EqualString(t *testing.T, res, resExpected string) {
	if res != resExpected {
		t.Error("Expected ", resExpected, " got ", res)
	}
}

// EqualUint16 use != operator to compare the arguments, then log it and fail.
func EqualUint16(t *testing.T, res, resExpected uint16) {
	if res != resExpected {
		t.Error("Expected ", resExpected, " got ", res)
	}
}

// NoErr use != operator to see if an error happened, then log it and fail.
func NoErr(t *testing.T, err error) {
	if err != nil {
		t.Error("Unexpected err ", err)
	}
}
