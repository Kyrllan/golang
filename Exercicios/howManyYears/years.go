package main

import "time"

func calculateMyAge(day, month, year int) int {
	birthDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	years := now.Year() - birthDate.Year()

	// Adjust age if birthday hasn't occurred this year
	if now.Month() < birthDate.Month() ||
		(now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		years--
	}

	return years
}

func main() {
	response := calculateMyAge(21, 9, 1989)
	println(response)
}
