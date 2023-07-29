// Test the calculateEaster function
package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCalculateEaster(t *testing.T) {

	var result bool
	var tests = []struct {
		year, month, day int
		want             bool
	}{
		{2023, 4, 11, false},
		{2023, 4, 10, false},
		{2023, 4, 9, true},
		{2023, 4, 8, false},
		{2023, 4, 7, false},

		{2024, 3, 31, true},
		{2022, 4, 17, true},
		{2021, 4, 4, true},
		{2020, 4, 12, true},
		{2019, 4, 21, true},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%d,%d,%d", tt.year, tt.month, tt.day)
		t.Run(testname, func(t *testing.T) {
			ans, _ := calculateEaster(tt.year)
			year, month, day := ans.Date()
			if year == tt.year && int(month) == tt.month && day == tt.day {
				result = bool(true)
			} else {
				result = bool(false)
			}

			if result != tt.want {
				t.Errorf("got %t, want %t", result, tt.want)
			}
		})
	}
}

// Test the isWeekday function
func TestIsWeekday(t *testing.T) {
	var tests = []struct {
		year, month, day int
		want             bool
	}{

		{2023, 4, 30, false}, // Sunday
		{2023, 4, 29, false}, // Saturday
		{2023, 4, 28, true},  // Friday
		{2023, 4, 27, true},  // Thursday
		{2023, 4, 26, true},  // Wednesday
		{2023, 4, 25, true},  // Tuesday
		{2023, 4, 24, true},  // Monday
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d,%d", tt.year, tt.month, tt.day)
		t.Run(testname, func(t *testing.T) {
			testdate := time.Date(tt.year, time.Month(tt.month), tt.day, 0, 0, 0, 0, time.UTC)
			result := isWeekday(testdate)
			if result != tt.want {
				t.Errorf("got %t, want %t", result, tt.want)
			}
		})
	}
}
