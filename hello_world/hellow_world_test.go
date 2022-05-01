package main

import (
	"testing"
)

func TestHello_v1(t *testing.T) {
	got := Hello_v1()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
func TestHello_v2(t *testing.T) {
	got := Hello_v2("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}


func TestHello_v2(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello_v2("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello_v2("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
*/

func TestHello_v2(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello_v2("Malli")
		want := "Hello, Malli"
		assertCorrectMsg(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello_v2("")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})
}

func TestHello_v3(t *testing.T) {
	t.Run("saying hello in English", func(t *testing.T) {
		got := Hello_v3("Malli", "")
		want := "Hello, Malli"
		assertCorrectMsg(t, got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello_v3("Malli", "spanish")
		want := "Hola, Malli"
		assertCorrectMsg(t, got, want)
	})
	t.Run("say hello in Hungary", func(t *testing.T) {
		got := Hello_v3("Malli", "hungary")
		want := "Szia, Malli"
		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "first test",
			args: "Malli",
			want: "Hello, Malli",
		},
		{
			name: "Second test",
			args: "Sakthi",
			want: "Hello, Sakthi",
		},
		{
			name: "Third test",
			args: "GokiMoki",
			want: "Hello, GokiMoki",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
