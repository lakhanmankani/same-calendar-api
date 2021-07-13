package samecalendar

import (
	"errors"
	"time"
)

func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	} else if year%100 != 0 {
		return true
	} else if year%400 != 0 {
		return false
	}
	return true
}

func yearStartDay(year int) time.Weekday {
	date := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	return date.Weekday()
}

func nextYearWithSameLeapness(year int) (nextYear int) {
	isLeap := isLeapYear(year)
	nextYear = year + 1
	for isLeapYear(nextYear) != isLeap {
		nextYear += 1
	}
	return nextYear
}

func nextYearWithSameStartDay(year int) (nextYear int) {
	startDay := yearStartDay(year)
	nextYear = year + 5

	for yearStartDay(nextYear) != startDay {
		nextYear += 1
	}
	return nextYear
}

func SameCalendar(year int, n int) ([]int, error) {
	if n < 0 {
		err := errors.New("n must be greater than 0")
		return nil, err
	}
	var years = []int{year}

	for len(years) < n {
		lastYear := years[len(years)-1]

		nextSameLeap := nextYearWithSameLeapness(lastYear)
		nextSameDay := nextYearWithSameStartDay(lastYear)
		for nextSameLeap != nextSameDay {
			if nextSameLeap > nextSameDay {
				nextSameDay = nextYearWithSameStartDay(nextSameDay)
			} else if nextSameDay > nextSameLeap {
				nextSameLeap = nextYearWithSameLeapness(nextSameLeap)
			}
		}
		years = append(years, nextSameLeap)
	}
	return years, nil
}
