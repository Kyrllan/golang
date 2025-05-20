package main

import "time"

func calculateMyAge(day, month, year int) int {
	timeNow := time.Now()
	yearNow := timeNow.Year()
	monthNow := int(timeNow.Month())
	dayNow := timeNow.Day()
	years := yearNow - year
	months := monthNow - month
	days := dayNow - day
	if months < 0 {
		years--
		months = 12 + months
	}
	if days < 0 {
		months--
		days = 30 + days
	}
	return years
}

func main() {
	response := calculateMyAge(21, 9, 1989)
	println(response)
}
