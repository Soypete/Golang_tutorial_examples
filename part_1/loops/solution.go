package loops

import (
	"time"
)

type Movie struct {
	Name          string
	Rating        string
	TimeInMinutes int
}

// CalculateWatchTime sums up the run time of all movies the argument "movies".
// this funciton is implemented with a range over a slice!
// func CaclulateWatchTime(movies []Movie) int {
// var totalTime int
// for _, movie := range movies {
// totalTime += movie.TimeInMinutes
// }
// return totalTime
// }

// CalculateWatchTime sums up the run time of all movies the argument "movies".
// this funciton is implemented with a c++ style for loop.
func CaclulateWatchTime(movies []Movie) int {
	var totalTime int
	for i := 0; i < len(movies); i++ {
		totalTime += movies[i].TimeInMinutes
	}
	return totalTime
}

// GetDaysLeftOfSchool calculates the number of days left in a given school year from todays date. It is written using
// a while loop syntax for loop. Arguments startMonth and isFullYear are used to determine
// if the school year is on a semester format or a year-long format. The endMonth and endDay
// arguments are used to end the loop on the correct day.
// this is not the most efficient way to calculate a number of days. This is written
// with the intent of giving an example of the while loop syntax. For simple solution use
// time.Sub()
func GetDaysLeftOfSchool(startMonth, endMonth, endDay int, isFullYear bool) int {
	// the the current year
	year, month, _ := time.Now().UTC().Date()
	endYear := year

	// check if we need to add a year to the end day
	if isFullYear && startMonth >= 7 && month >= 7 {
		endYear++
	}

	// start counter
	dateTracker := time.Now().UTC()
	var daysLeftOfSchool int

	// this continues to run until the endDay and the endMonth match the days we are adding to
	for dateTracker.Day() != endDay || dateTracker.Month() != time.Month(endMonth) {

		// add one day to dateTracker
		dateTracker = dateTracker.AddDate(0, 0, 1)

		// add one day to days left
		daysLeftOfSchool++
	}

	// return days left
	return daysLeftOfSchool
}
