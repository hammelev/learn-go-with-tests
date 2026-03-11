package main

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeatedString := Repeat("hi", 3)
	fmt.Print(repeatedString)
	// Output: hihihi
}
