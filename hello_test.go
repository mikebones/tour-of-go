package main

import "testing"

func TestMain(t *testing.T) {
	got,got2 := swap("x", "y")
	want, want2 := "y" , "x"
	if got != want || got2 != want2 {
		t.Errorf("Got %q want %q, got2 %q, want %q", got, want, got2, want2)
	}
}
