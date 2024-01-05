package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("", "jhon")
		want := "Hello, jhon"

		AssetCorrectMessage(t, got, want)

	})

	t.Run(" saying hello with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		AssetCorrectMessage(t, got, want)
	})

	t.Run("sayng hello in spanish", func(t *testing.T) {
		got := Hello("spanish", "seba")
		want := "Hola, seba"
		AssetCorrectMessage(t, got, want)
	})

	t.Run("sayng hello in french", func(t *testing.T) {
		got := Hello("french", "seba")
		want := "Bonjour, seba"
		AssetCorrectMessage(t, got, want)
	})
}

func AssetCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
