package gosort

import (
	"reflect"
	"testing"
)

func fail[T any](t *testing.T, msg string, got T, expected T) {
	t.Helper()
	t.Errorf("Test failed. %s. Got: %v, expected: %v", msg, got, expected)
}

func assertTrue(t *testing.T, v bool) {
	t.Helper()

	if v {
		return
	}

	fail(t, "Value was not true", v, true)
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err == nil || reflect.ValueOf(err).IsNil() {
		return
	}
	fail(t, "Expected no error", err, nil)
}

func assertIsNilSlice[T any](t *testing.T, got []T) {
	t.Helper()

	if cap(got) != 0 {
		fail(t, "Capacity was not 0", cap(got), 0)
	}
	if len(got) != 0 {
		fail(t, "Length was not 0", len(got), 0)
	}
}

func assertSlicesEqual[T comparable](t *testing.T, got, expected []T) {
	t.Helper()

	if SlicesAreEqual(got, expected) {
		return
	}

	fail(t, "Not equal", got, expected)
}

func assertSlicesNotEqual[T comparable](t *testing.T, got, expected []T) {
	t.Helper()

	if !SlicesAreEqual(got, expected) {
		return
	}

	fail(t, "Equal", got, expected)
}
