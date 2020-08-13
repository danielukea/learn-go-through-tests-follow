package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {

	t.Run("test racer", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(10 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := ConfigurableRacer(slowURL, fastURL, 25*time.Millisecond)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("test server time out", func(t *testing.T) {
		serverA := makeDelayedServer(21 * time.Millisecond)
		serverB := makeDelayedServer(20 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 10*time.Millisecond)

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
