package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		// Arrange
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL

		// Act
		got, err := Racer(slowURL, fastURL)

		// Assert
		if err != nil {
			t.Fatalf("did not expect an error but gone one %v", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("return an error if no server responds within 10s", func(t *testing.T) {
		// Arrange
		serverA := makeDelayedServer(25 * time.Millisecond)

		defer serverA.Close()

		// Act
		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 20*time.Millisecond)

		// Assert
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
