package main

// go get github.com/vjeantet/eastertime

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
