package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectmEssage := func(t *testing.T, got, want string) {

		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}

	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := englishHelloPrefix + "Chris"
		assertCorrectmEssage(t, got, want)

	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := englishHelloPrefix + "World"

		assertCorrectmEssage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := spanishHelloPrefix + "Elodie"
		assertCorrectmEssage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Moe", "French")
		want := frenchHelloPrefix + "Moe"
		assertCorrectmEssage(t, got, want)
	})
}
