package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	cases := map[string]int{
		"III":     3,
		"IV":      4,
		"LVIII":   58,
		"MCMXCIV": 1994,
	}

	for input, want := range cases {
		result := RomanToInt(input)
		if result != want {
			t.Fatalf(`romanToInt("%s") = %d, wanted %d`, input, result, want)
		}
	}
}

// if a character is not mapped
// a zero value of int ( 0 ) will be used
func TestInvalidChar(t *testing.T) {
	cases := map[string]int{
		"Z":   0,
		"ZI":  1,
		"IZ":  1,
		"IZI": 2,
	}

	for input, want := range cases {
		result := RomanToInt(input)
		if result != want {
			t.Fatalf(`romanToInt("%s") = %d, wanted %d`, input, result, want)
		}
	}

}

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomanToInt("MCMXCIV")
	}
}
