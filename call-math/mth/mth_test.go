package mth

import (
	"testing"
)

// TDD and BDD
// TDD- testing code first and behaviour driven development ginkgo
func TestAdd(t *testing.T) {
	got := Add(5, 4)
	want := 9

	if got != want {
		t.Errorf("test add failed got: %#v %#v", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(5, 4)
	want := 1

	if got != want {
		t.Errorf("test add failed got: %#v %#v", got, want)
	}
}
