package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("with structs, pointers, slices, and arrays", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         any
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Chris", "Vienna"},
				[]string{"Chris", "Vienna"},
			},
			{
				"struct with non string field",
				struct {
					Name string
					Age  int
				}{"Chirs", 30},
				[]string{"Chirs"},
			},
			{
				"nested fields",
				Person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"pointers to things",
				&Person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"slices",
				[]Profile{
					Profile{39, "Vienna"},
					Profile{33, "London"},
				},
				[]string{"Vienna", "London"},
			},
			{
				"arrays",
				[2]Profile{
					Profile{39, "Vienna"},
					Profile{33, "London"},
				},
				[]string{"Vienna", "London"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				want := test.ExpectedCalls

				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			})
		}
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Vienna"}
			aChannel <- Profile{28, "Copenhagen"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Vienna", "Copenhagen"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Vienna"}, Profile{28, "Copenhagen"}
		}

		var got []string
		want := []string{"Vienna", "Copenhagen"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
