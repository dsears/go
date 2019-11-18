package dsears

import "fmt"
import "math"

// DaysToShortUnit : Take an integer number of days and format it using the closest time unit (days, weeks, months or years).
// Values are rounded to one decimal point, and we use the largest unit with a value >= 1.0.
func DaysToShortUnit(originalDays int) string {
	atLeast1_0 := func(f float64) bool {
		return int(math.Round(f*10.0)) >= 10
	}
	days := float64(originalDays)
	year := 365.25
	years := days / year
	if atLeast1_0(years) {
		return fmt.Sprintf("%0.1f years", years)
	}
	month := year / 12.0
	months := days / month
	if atLeast1_0(months) {
		return fmt.Sprintf("%0.1f months", months)
	}
	weeks := days / 7.0
	if atLeast1_0(weeks) {
		return fmt.Sprintf("%0.1f weeks", weeks)
	}
	return fmt.Sprintf("%0.1f days", days)
}
