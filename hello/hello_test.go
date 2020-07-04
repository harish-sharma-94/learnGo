package main

import "testing"

func TestHello(t *testing.T) {

	assetCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("saying hello to poeple", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assetCorrectMessage(t, got, want)
	})

	t.Run("saying hello to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assetCorrectMessage(t, got, want)
	})
}
