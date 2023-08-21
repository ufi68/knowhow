package main

import (
	"fmt"
	"time"

	cal "github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

func main() {
	c := cal.NewBusinessCalendar()
	c.Name = "T2S"
	c.Description = "Default company calendar"
	// https://www.ecb.europa.eu/paym/groups/shared/docs/35c5c-ami-pay-2017-12-06-joint-ami-pay-ami-seco-item-3-t2s-on-target2-closing-days.pdf

	// add holidays that the business observes
	c.AddHoliday(
		aa.GoodFriday,
		aa.EasterMonday,
		aa.WorkersDay, // 1th of May
	)

	// change the default of a Mon - Fri, 9am-5pm work week
	// t := time.Now()
	// t := time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t := time.Date(2024, time.Month(12), 24, 0, 0, 0, 0, time.UTC)

	// run different batch processing jobs based on the day

	fmt.Println(c.IsWorkday(t))
	fmt.Println(cal.IsWeekend(t))

	if c.IsWorkday(t) {

		// create daily activity reports
	}

	if cal.IsWeekend(t) {
		// validate employee time submissions
	}

	if c.WorkdaysRemain(t) == 10 {
		// 10 business days before the end of month
		// send account statements to customers
	}

	if c.WorkdaysRemain(t) == 0 {
		// last business day of the month
		// execute auto billing transfers
	}

	// determine the number of working hours left in the current month
	nextMonth := cal.DayStart(cal.MonthStart(t.AddDate(0, 1, 0)))
	hoursLeft := c.WorkHoursInRange(t, nextMonth)

	// check if there are any tasks for this month that are in danger of missing their deadline
	pendingTasks := []struct{ pendingHours time.Duration }{{pendingHours: 32}} // assumed to be fetched from a DB or API
	for _, task := range pendingTasks {
		if hoursLeft < task.pendingHours {
			// send alert to management
		}
	}
}
