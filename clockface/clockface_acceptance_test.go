package clockface_test

import (
	"testing"
	"time"

	"github.com/learn-go-luke/clockface"
)

func TestSecondHand(t *testing.T) {
	t.Run("test second hand at midnight", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		want := clockface.Point{X: 150.0, Y: 150 - 90}
		got := clockface.SecondHand(tm)

		if got != want {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})

	// t.Run("test second hand at 30 seconds", func(t *testing.T) {
	// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	// 	want := clockface.Point{X: 150.0, Y: 150 + 90}
	// 	got := clockface.SecondHand(tm)

	// 	if got != want {
	// 		t.Errorf("Got %v, wanted %v", got, want)
	// 	}
	// })

}
