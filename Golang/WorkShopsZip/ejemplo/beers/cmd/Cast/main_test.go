package main

import (
	"testing"
)

//TestNumToString Convierte entero a string
func TestNumToString(t *testing.T) {
	got := NumToString(8)
	want := "8"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
