package samecalendar

import (
	"errors"
	"time"
)

var ErrorNegativeYear = errors.New("year must not be negative")
var ErrorNegativeN = errors.New("n must be greater than 0")

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

func previousYearWithSameLeapness(year int) (nextYear int) {
	isLeap := isLeapYear(year)
	nextYear = year - 1
	for isLeapYear(nextYear) != isLeap {
		nextYear -= 1
	}
	return nextYear
}

func previousYearWithSameStartDay(year int) (nextYear int) {
	startDay := yearStartDay(year)
	nextYear = year - 5

	for yearStartDay(nextYear) != startDay {
		nextYear -= 1
	}
	return nextYear
}

func forwardSameCalendar(year int, n int) (years []int, err error) {
	if year < 0 {
		return nil, ErrorNegativeYear
	}
	if n < 0 {
		return nil, ErrorNegativeN
	}
	if n == 0 {
		return []int{}, nil
	}
	years = []int{year}

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

func backwardSameCalendar(year int, n int) (years []int, err error) {
	if year < 0 {
		return nil, ErrorNegativeYear
	}
	if n < 0 {
		return nil, ErrorNegativeN
	}
	if n == 0 {
		return []int{}, nil
	}
	years = []int{year}

	for len(years) < n {
		lastYear := years[len(years)-1]

		prevSameLeap := previousYearWithSameLeapness(lastYear)
		prevSameDay := previousYearWithSameStartDay(lastYear)
		for prevSameLeap != prevSameDay {
			if prevSameLeap < prevSameDay {
				prevSameDay = previousYearWithSameStartDay(prevSameDay)
			} else if prevSameDay < prevSameLeap {
				prevSameLeap = previousYearWithSameLeapness(prevSameLeap)
			}
		}
		if prevSameLeap < 0 {
			break
		}
		years = append(years, prevSameLeap)
	}
	return years, nil
}

func SameCalendar(year int, n int, forward bool) (years []int, err error) {
	if forward {
		return forwardSameCalendar(year, n)
	}
	return backwardSameCalendar(year, n)
}
