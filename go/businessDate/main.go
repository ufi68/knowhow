package main

// go get github.com/vjeantet/eastertime

// Alternative - use
// https://github.com/rickar/cal

import (
	"fmt"
	"time"

	eastertime "github.com/vjeantet/eastertime"
)

func main() {

	// Date to calculate
	checkDateTime := time.Now()
	if isEaster(checkDateTime) {
		fmt.Println("Ostern ;)")
	}
	fmt.Println("today is a Weekday", checkDateTime, isWeekday(time.Now()))
}

func calculateEaster(year int) (time.Time, error) {

	return eastertime.CatholicByYear(year)
}

func isEaster(date time.Time) bool {

	year, month, day := date.Date()

	easternDate, _ := calculateEaster(year)
	_, easternDateMonth, easternDateDay := easternDate.Date()

	if month == easternDateMonth && day == easternDateDay {
		return true
	}

	return false
}

// return false if the weekday in the time object is a Saturday or an Sunday, otherwise true
func isWeekday(date time.Time) bool {
	if date.Weekday() != time.Saturday && date.Weekday() != time.Sunday {
		return true
	}
	return false
}

// Create a function that takes a date and returns true if the date is a weekend, false otherwise
// Hint: use the time.Weekday() function
func isWeekend(date time.Time) bool {
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return true
	}
	return false
}
