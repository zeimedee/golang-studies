package main

import (
	"fmt"
	"time"
)

func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
	panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	today := time.Now()
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	tt := t.Before(today)
	if tt {
		return true
	}
	return false
	panic("Please implement the HasPassed function")
}

func main() {
	// a := HasPassed("October 3, 2019 20:32:00")
	// b := HasPassed("January 28, 1974 2:02:02")
	// c := HasPassed("December 9, 2112 11:59:59")
	// fmt.Println("a =", a)
	// fmt.Println("b =", b)
	// fmt.Println("c =", c)

	// t, _ := time.Parse("Monday, January 2, 2006 15:04:05", "Friday, March 8, 1974 12:02:02")
	// fmt.Println(t)
	// fmt.Println(t.Hour())

	// date := Schedule("7/25/2019 13:45:00")
	// fmt.Println(date)
	// d := date.Weekday()
	// dy := date.Day()
	// m := date.Month()
	// y := date.Year()
	// h := date.Hour()
	// min := date.Minute()

	// f := fmt.Sprintf("You have an appointment on %d, %s %d, %d, at %d:%d.", d, m, dy, y, h, min)
	// fmt.Println(f)

	t := time.Now()
	fmt.Println(t)

}
