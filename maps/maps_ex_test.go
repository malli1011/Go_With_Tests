package maps

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("With dictonary as argument", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}
		Init()
		got := Search(dictionary, "test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

	t.Run("With dictonary nil", func(t *testing.T) {
		Init()
		got := Search(nil, "test")
		want := "default test value"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

}

func TestDictonary(t *testing.T) {
	dictionary := Dictionary{"test": "validate the outcome", "work": "collection of letters"}

	t.Run("Expect no error", func(t *testing.T) {
		got, err := dictionary.Search("test")
		expected := "validate the outcome"

		if err != nil {
			t.Fatal("Error not expected here!")
		}

		if got != expected {
			t.Errorf("expected %s, but got %s", expected, got)
		}
	})

	t.Run("Expect Error", func(t *testing.T) {
		_, err := dictionary.Search("noword")

		if err == nil {
			t.Fatal("Expected to throw error")
		}

		if err.Error() != ErrNotFound.Error() {
			t.Error("Error message not matching for missing word")
		}
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	t.Run("Should add new word to dictionary", func(t *testing.T) {
		err := dictionary.Add(word, definition)
		if err != nil {
			t.Fatal("Error not expected here")
		}
		assertDefinition(t, dictionary, word, definition)

	})

	t.Run("Should throw error", func(t *testing.T) {
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExists.Error())
	})

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got error, expected string) {
	t.Helper()

	if got.Error() != expected {
		t.Errorf("got %q want %q", got.Error(), expected)
	}

}

func TestDictionary_update(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		if err != nil {
			t.Fatal("Error not expected here")
		}

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrNotFound.Error())
	})
}

func TestDictionary_print(t *testing.T) {
	dictionary := Dictionary{"test": "Test value", "test1": "Test value2"}
	for key, val := range dictionary {
		fmt.Printf("%v \t %v\n", key, val)
	}

	keys := make([]string, 0, len(dictionary))
	for k := range dictionary {
		keys = append(keys, k)
	}

	for _, key := range keys {
		fmt.Println(key)
	}
}
