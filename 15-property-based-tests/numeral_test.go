package numeral

import (
	"fmt"
	"strconv"
	"testing"
	"testing/quick"
)

type romanNumeralTest struct {
	Description string
	Arabic      uint16
	Roman       string
}

func testCaseGenerator(arabic uint16, roman string) romanNumeralTest {
	return romanNumeralTest{
		Description: strconv.FormatUint(uint64(arabic), 10) + " -> " + roman,
		Arabic:      arabic,
		Roman:       roman,
	}
}

var cases = []romanNumeralTest{
	testCaseGenerator(1, "I"),
	testCaseGenerator(2, "II"),
	testCaseGenerator(3, "III"),
	testCaseGenerator(4, "IV"),
	testCaseGenerator(5, "V"),
	testCaseGenerator(6, "VI"),
	testCaseGenerator(7, "VII"),
	testCaseGenerator(8, "VIII"),
	testCaseGenerator(9, "IX"),
	testCaseGenerator(10, "X"),
	testCaseGenerator(14, "XIV"),
	testCaseGenerator(18, "XVIII"),
	testCaseGenerator(20, "XX"),
	testCaseGenerator(39, "XXXIX"),
	testCaseGenerator(40, "XL"),
	testCaseGenerator(47, "XLVII"),
	testCaseGenerator(49, "XLIX"),
	testCaseGenerator(50, "L"),
	testCaseGenerator(90, "XC"),
	testCaseGenerator(100, "C"),
	testCaseGenerator(400, "CD"),
	testCaseGenerator(500, "D"),
	testCaseGenerator(798, "DCCXCVIII"),
	testCaseGenerator(900, "CM"),
	testCaseGenerator(1006, "MVI"),
	testCaseGenerator(1984, "MCMLXXXIV"),
	testCaseGenerator(2014, "MMXIV"),
	testCaseGenerator(3999, "MMMCMXCIX"),
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Roman

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			want := test.Arabic

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			t.Log("testing", arabic)
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
