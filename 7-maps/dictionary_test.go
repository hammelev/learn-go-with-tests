package maps

import (
	"testing"
)

var testKey, testVal = "test", "this is just a test"

func TestSearch(t *testing.T) {
	// Arrange
	dictionary := Dictionary{testKey: testVal}

	t.Run("Search successfull", func(t *testing.T) {
		assertDefinition(t, dictionary, testKey, testVal)
	})

	t.Run("Search for non-existent word", func(t *testing.T) {
		// Arrange
		nonExistentWord := "non-existent"
		want := ErrNotFound

		// Act
		_, err := dictionary.Search(nonExistentWord)

		// Assert
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {

	t.Run("Add successfull", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{}

		// Act
		err := dictionary.Add(testKey, testVal)

		// Assert
		assertError(t, err, nil)
		assertDefinition(t, dictionary, testKey, testVal)
	})

	t.Run("Add duplicate", func(t *testing.T) {
		// Arrange
		newDefinition := "new value"
		dictionary := Dictionary{testKey: testVal}
		want := ErrWordExists

		// Act
		err := dictionary.Add(testKey, newDefinition)

		// Assert
		assertError(t, err, want)
		assertDefinition(t, dictionary, testKey, testVal)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update exisiting word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{testKey: testVal}
		newDefinition := "new value"

		// Act
		err := dictionary.Update(testKey, newDefinition)

		// Assert
		assertError(t, err, nil)
		assertDefinition(t, dictionary, testKey, newDefinition)
	})

	t.Run("Update non-exisiting word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{}

		// Act
		err := dictionary.Update(testKey, testVal)

		// Assert
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete existing word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{testKey: testVal}

		// Act
		err := dictionary.Delete(testKey)

		// Assert
		assertError(t, err, nil)

		// Act
		_, err = dictionary.Search(testKey)

		// Assert
		assertError(t, err, ErrNotFound)
	})

	t.Run("Delete non-existing word", func(t *testing.T) {
		// Arrange
		dictionary := Dictionary{}

		// Act
		err := dictionary.Delete(testKey)

		// Assert
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("expected no error", err)
	}
	assertStrings(t, got, definition)
}
