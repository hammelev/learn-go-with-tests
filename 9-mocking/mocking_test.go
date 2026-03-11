package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

const write = "write"
const sleep = "sleep"

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		// Arrange
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		// Act
		Countdown(buffer, spySleeper)
		got := buffer.String()

		// Assert
		want := "3\n2\n1\n" + finalWord

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		// Arrange
		spySleepPrinter := &SpyCountdownOperations{}

		//Act
		Countdown(spySleepPrinter, spySleepPrinter)

		// Assert
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(spySleepPrinter.Calls, want) {
			t.Errorf("got %v want %v", spySleepPrinter.Calls, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	// Arrange
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}

	// Act
	sleeper.Sleep()

	// Assert
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}
