package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Brenden")
	want := "Hello Brenden"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
