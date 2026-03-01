package helloWorld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("", "Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Spanish", "Elodie")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("French", "Napoleon")
		want := "Bonjour, Napoleon"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Danish", func(t *testing.T) {
		got := Hello("Danish", "Mikkel")
		want := "Hej, Mikkel"
		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
