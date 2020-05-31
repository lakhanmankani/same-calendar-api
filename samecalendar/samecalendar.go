package samecalendar

import (
	"errors"
	"fmt"
	"time"
)

func isLeapYear(year int) bool {
	if year % 4 != 0 {
		return false
	} else if year % 100 != 0 {
		return true
	} else if year % 400 != 0 {
		return false
	}
	return true
}

func yearStartDay(year int) time.Weekday {
	date := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	return date.Weekday()
}


func yearsWithSameLeapness(year int, n int) []int {
	var years []int
	yearLeapness := isLeapYear(year)

	for len(years) < n {
		if isLeapYear(year) == yearLeapness {
			years = append(years, year)
		}
		year += 1
	}
	return years
}


func yearsWithSameStartDay(year int, n int) []int {
	var years []int
	startDay := yearStartDay(year)

	for len(years) < n {
		if yearStartDay(year) == startDay {
			years = append(years, year)
		}
		year += 1
	}
	return years
}


func intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}


func SameCalendar(year int, n int) ([]int, error) {
	if n == 1 {
		return []int{year}, nil
	}
	if n < 0 {
		err := errors.New("n must be greater than 0")
		return nil, err
	}
	var years []int

	for len(years) < n {
		var lastYear int
		if len(years) == 0 {
			lastYear = year
		} else {
			lastYear = years[len(years)-1]
		}

		sameLeapness := yearsWithSameLeapness(lastYear, 10*n)
		//fmt.Println(sameLeapness)
		sameStartDate := yearsWithSameStartDay(lastYear, 10*n)
		//fmt.Println(sameStartDate)
		yearIntersection := intersection(sameLeapness, sameStartDate)
		years = append(years, yearIntersection...)
	}
	fmt.Println(years)
	return years[:n], nil
}