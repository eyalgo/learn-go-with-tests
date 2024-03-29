package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
	}
	
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Eyal", "")
    	want := "Hello, Eyal"

    	assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
    	want := "Hello, World"

    	assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })
    
	t.Run("in French", func(t *testing.T) {
        got := Hello("Merrie", "French")
        want := "Bonjour, Merrie"
        assertCorrectMessage(t, got, want)
    })
}
