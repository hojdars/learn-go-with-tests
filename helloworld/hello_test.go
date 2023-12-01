package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("just say hello world if no name is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
