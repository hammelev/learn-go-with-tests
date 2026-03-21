package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			want := test.angle
			got := secondsInRadians(test.time)

			if !roughlyEqualFloat64(got, want) {
				t.Fatalf("wanted %v radians, but got %v", want, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			want := test.point
			got := secondHandPoint(test.time)

			if !roughlyEqualPoint(got, want) {
				t.Fatalf("wanted %v point, but got %v", want, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time time.Time
		angle float64
	} {
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.angle
			got := minutesInRadians(c.time)

			if !roughlyEqualFloat64(got, want) {
				t.Fatalf("Wanted %v radians, but got %v", want, got)
			}
		})
	}
}

func TestMinutesHandPoint(t *testing.T) {
	cases := []struct {
		time time.Time
		point Point
	} {
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.point
			got := minuteHandPoint(c.time)
			if !roughlyEqualPoint(got, want) {
				t.Fatalf("wanted %v point, but got %v", want, got)
			}

		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time time.Time
		angle float64
	} {
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.angle
			got := hoursInRadians(c.time)

			if !roughlyEqualFloat64(got, want) {
				t.Fatalf("Wanted %v radians, but got %v", want, got)
			}
		})
	}
}

func TestHoursHandPoint(t *testing.T) {
	cases := []struct {
		time time.Time
		point Point
	} {
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.point
			got := hourHandPoint(c.time)
			if !roughlyEqualPoint(got, want) {
				t.Fatalf("wanted %v point, but got %v", want, got)
			}

		})
	}
}


func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(312, time.October, 28, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
