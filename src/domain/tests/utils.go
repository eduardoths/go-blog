package tests

import "testing"

const AssertFailed string = "Assert failed: Expected value = %v, actual value = %v"

func AssertEquals(t testing.TB, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf(AssertFailed, expected, actual)
	}
}